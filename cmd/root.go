/*
Copyright © 2024 @0xNader
License MIT
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BasePaint/bpverify/pkg/img"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "bpverify",
    Short: "Verify the final image for a day on BasePaint.",
    Long: `This command allows you to generate and verify the final image for a specific BasePaind day.
    This command fetches all paint events for the specified day from the blockchain, processes them, and creates a PNG image representing the final state of the BasePaint canvas.

    Usage:
    bpverify [flags]

    Flags:
    -r, --rpc string       RPC URL for the Base L2 node (required)
    -d, --day int          Day of BasePaint to verify (required)
    -p, --path string      Path to save the generated image (default is your Desktop)

    Examples:

    Verify BasePaint Day #5
    bpverify --rpc https://base-mainnet.g.alchemy.com/v2/API_KEY --day 5

    Verify BasePaint Day #10 and save img to a specific path
    bpverify -r https://base-mainnet.g.alchemy.com/v2/API_KEY -d 10 -p /path/to/save/BasePaint#10.png`,
    Run: func(cmd *cobra.Command, args []string) {
        err := img.CreateImage(rpc, day, path)
        if err != nil {
            fmt.Printf("Error getting events: %v\n", err)
        }
    },
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        if !cmd.Flags().Changed("outputPath") {
            desktopPath, err := getDesktopPath()
            if err != nil {
                desktopPath = "."
            }

            fileName := fmt.Sprintf("BP Verify Day#%d.png", day)
            imgFilePath := filepath.Join(desktopPath, fileName)

            cmd.Flags().Set("outputPath", imgFilePath)
        }
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	rpc string;
	day int;
	path string;
)

func init() {
	desktopPath, err := getDesktopPath()
    if err != nil {
        desktopPath = "."
    }

	fileName := "BP Verify.png"
	imgFilePath := filepath.Join(desktopPath, fileName)


    rootCmd.Flags().StringVarP(&rpc, "rpc", "r", "https://mainnet.infura.io/v3/YOUR-PROJECT-ID", "RPC URL for Base L2 node")
    rootCmd.Flags().IntVarP(&day, "day", "d", 1, "Day of BasePaint to verify")
    rootCmd.Flags().StringVarP(&path, "outputPath", "p", imgFilePath, "Path to where to save image to")

    rootCmd.MarkFlagRequired("rpc")
    rootCmd.MarkFlagRequired("day")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getDesktopPath() (string, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(homeDir, "Desktop"), nil
}
