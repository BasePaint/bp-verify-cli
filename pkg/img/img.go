package img

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"os"
	"strconv"

	"github.com/BasePaint/bpverify/pkg/events"
)

type Config struct {
	Theme   string   `json:"theme"`
	Palette []string `json:"palette"`
	Size    int      `json:"size"`
}

const FinalImageSize = 1024

func CreateImage(rpcURL string, day int, outputPath string) error {
	apiURL := fmt.Sprintf("https://basepaint.xyz/api/theme/%d", day)
	config, err := fetchConfig(apiURL)
	if err != nil {
		return fmt.Errorf("error fetching config: %v", err)
	}

	initialCanvas := image.NewRGBA(image.Rect(0, 0, config.Size, config.Size))

	pixelsData, err := events.GetEvents(rpcURL, day)
	if err != nil {
		return fmt.Errorf("error getting events: %v", err)
	}

	colors, err := hexPaletteToRGBA(config.Palette)
	if err != nil {
		return fmt.Errorf("error converting palette: %v", err)
	}

	for _, pixels := range pixelsData {
		applyPixels(initialCanvas, pixels, colors, config.Size)
	}

	finalImg := image.NewRGBA(image.Rect(0, 0, FinalImageSize, FinalImageSize))
	scale := FinalImageSize / config.Size

	for y := 0; y < config.Size; y++ {
		for x := 0; x < config.Size; x++ {
			color := initialCanvas.At(x, y)
			rect := image.Rect(x*scale, y*scale, (x+1)*scale, (y+1)*scale)
			draw.Draw(finalImg, rect, &image.Uniform{color}, image.Point{}, draw.Src)
		}
	}

	if err := saveImage(finalImg, outputPath); err != nil {
		return fmt.Errorf("error saving image: %v", err)
	}

	fmt.Printf("Image created successfully: %s\n", outputPath)
	fmt.Printf("Theme: %s\n", config.Theme)
	fmt.Printf("Original Canvas Size: %d\n", config.Size)
	fmt.Printf("Final Image Size: %d\n", FinalImageSize)

	return nil
}

func fetchConfig(url string) (*Config, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var config Config
	if err := json.NewDecoder(resp.Body).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func hexPaletteToRGBA(hexColors []string) ([]color.RGBA, error) {
	rgbaColors := make([]color.RGBA, len(hexColors))
	for i, hex := range hexColors {
		r, g, b, err := hexToRGB(hex)
		if err != nil {
			return nil, err
		}
		rgbaColors[i] = color.RGBA{R: r, G: g, B: b, A: 255}
	}
	return rgbaColors, nil
}

func hexToRGB(hex string) (uint8, uint8, uint8, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}
	var r, g, b uint8
	_, err := fmt.Sscanf(hex[1:], "%02x%02x%02x", &r, &g, &b)
	return r, g, b, err
}

func applyPixels(img *image.RGBA, pixels []byte, colors []color.RGBA, size int) {
    hexString := hex.EncodeToString(pixels)

	// Instruction for 1 pixel is 6 chars long
	// x, y, color index
    for i := 0; i < len(hexString); i += 6 {
        if i+6 > len(hexString) {
            break
        }

        x, _ := strconv.ParseUint(hexString[i:i+2], 16, 8)
        y, _ := strconv.ParseUint(hexString[i+2:i+4], 16, 8)
        colorIndex, _ := strconv.ParseUint(hexString[i+4:i+6], 16, 8)

        if int(x) < size && int(y) < size && int(colorIndex) < len(colors) {
            img.Set(int(x), int(y), colors[int(colorIndex)])
        }
    }
}

func saveImage(img *image.RGBA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()
	return png.Encode(f, img)
}