# Image Analyzer

A lightweight command-line image analysis tool written in Go that extracts image metadata and inspects pixel color values.

## Features

- Extract image metadata from supported image files
- Display camera and EXIF information
- Read image dimensions
- Extract GPS coordinates when available
- Inspect RGB values of a pixel
- Interactive CLI interface
- Cross-platform support

## Project Structure

```text
image-analyzer/
├── analyze-metadata/
│   └── analyzer.go
├── cli/
│   ├── image-processing.go
│   ├── reader.go
│   └── start.go
├── main.go
├── go.mod
└── go.sum
```

## Requirements

- Go 1.20 or newer

## Installation

Clone the repository:

```bash
git clone https://github.com/your-username/image-analyzer.git
cd image-analyzer
```

Install dependencies:

```bash
go mod tidy
```

Build the application:

```bash
go build -o image-analyzer
```

## Usage

Run the application:

```bash
go run .
```

or

```bash
./image-analyzer
```

## Commands

### Show Help

```text
help
```

Displays available commands.

### Read Metadata

```text
metadata <image-path>
```

Example:

```text
metadata photo.jpg
```

Output may include:

```text
Camera Make: Canon
Camera Model: EOS 80D
Date/Time: 2025-05-20 14:32:10
Focal Length: 50.0
ISO: 100
Width: 6000
Height: 4000
Latitude: 19.0760
Longitude: 72.8777
GPS: true
```

### Read RGB Information

```text
rgb <image-path>
```

Example:

```text
rgb image.png
```

Output:

```text
Pixel: R=65535 G=32768 B=16384 A=65535
```

> Note: The current implementation reads the pixel located at coordinate `(100,100)`.

### Clear Screen

```text
clear
```

Clears the terminal.

### Exit

```text
exit
```

Terminates the application.

## Dependencies

- github.com/FlavioCFOliveira/GoMetadata

## Notes

- Metadata availability depends on the image containing EXIF information.
- GPS coordinates are only displayed when location metadata exists.
- Unsupported image formats may result in decoding errors.
- RGB inspection currently uses a fixed pixel position.

## Future Improvements

- Custom pixel coordinate selection
- Batch image processing
- Metadata export (JSON/CSV)
- Histogram generation
- Advanced color analysis
- Additional image format support

## License

MIT License
