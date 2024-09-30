# BasePaint Verify CLI Tool

BasePaint Verify is a command-line tool for generating and verifying images from a specific day on BasePaint. It fetches paint events from the blockchain, processes them, and creates a PNG image representing the final state of the BasePaint canvas.

## Prerequisites

Before you begin, ensure you have the following installed:

- Go 1.16 or later (Easiest way to install Go is to use [Webi](https://webinstall.dev/golang/))
- Access to a Base RPC endpoint ([Alchemy](https://www.alchemy.com), [Infura](https://www.infura.io), etc)

## Installation

1. Ensure you have Go installed on your system (version 1.16 or later). You can download it from [https://golang.org/dl/](https://golang.org/dl/)

2. Clone the repository:

   ```
   git clone https://github.com/yourusername/bpverify.git
   cd bpverify
   ```

3. Download the required dependencies:

   ```
   go mod download
   ```

4. Build the CLI tool:

   ```
   go build -o bpverify
   ```

5. (Optional) Add the tool to your PATH for easy access:
   ```
   sudo mv bpverify /usr/local/bin/
   ```

Note: If you're using Windows, the commands may differ slightly:

- Instead of `sudo mv`, you might add the directory containing `bpverify.exe` to your PATH environment variable.
- The executable will be named `bpverify.exe`.

## Usage

The basic syntax for using the tool is:

```
bpverify [flags]
```

### Flags

- `-r, --rpc`: RPC URL for the Base node (required)
- `-d, --day`: BasePaint Day to verify (required)
- `-o, --output`: Path to save the generated image (defaults to Desktop)

### Example

To generate an image for day 5 of BasePaint:

```
bpverify -r https://base-mainnet.g.alchemy.com/v2/API_KEY -d 5
```

## Configuration

The tool automatically fetches the configuration (theme, color palette, canvas size) for each day from the BasePaint API. No additional configuration is required.

## Troubleshooting

If you encounter any issues:

1. Ensure you have a valid RPC URL for a Base node.
2. Check that the specified day is valid.
3. Verify that you have write permissions in the directory where you're saving the output image.

## License

This project is licensed under the MIT License
