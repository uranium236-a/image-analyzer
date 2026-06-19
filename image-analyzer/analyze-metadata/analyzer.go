package analyzemetadata

import (
	"fmt"
	"image"
	"os"

	gometadata "github.com/FlavioCFOliveira/GoMetadata"
)

type ImageMetadata struct {
	CameraMake   string
	CameraModel  string
	DateTime     string
	FocalLength  string
	ISO          string
	Width        int
	Height       int
	GPSLatitude  float64
	GPSLongitude float64
	HasGPS       bool
	Error        error
}

func Analyze(filepath string) ImageMetadata {
	var metadata ImageMetadata

	// Read metadata from the file
	m, err := gometadata.ReadFile(filepath)
	if err != nil {
		metadata.Error = err
		return metadata
	}

	// Camera Make - returns string directly
	if makeVal := m.Make(); makeVal != "" {
		metadata.CameraMake = makeVal
	}

	// Camera Model - returns string directly
	if modelVal := m.CameraModel(); modelVal != "" {
		metadata.CameraModel = modelVal
	}

	// DateTime - returns (time.Time, bool)
	if dateTime, ok := m.DateTimeOriginal(); ok {
		metadata.DateTime = dateTime.String()
	} else if dateTime, ok := m.DateTime(); ok {
		metadata.DateTime = dateTime.String()
	}

	// Focal Length - returns (float64, bool) as a single float value
	if focalLen, ok := m.FocalLength(); ok {
		metadata.FocalLength = fmt.Sprintf("%.1f", focalLen)
	}

	// ISO - returns (uint, bool)
	if iso, ok := m.ISO(); ok {
		metadata.ISO = fmt.Sprintf("%d", iso)
	}

	// Image Dimensions - use Format() or EXIF pixel dimensions
	// Note: GoMetadata doesn't have a direct ImageDimensions method.
	// Use EXIF pixel dimensions instead.
	if metadata.Width == 0 || metadata.Height == 0 {
		if imgFile, err := os.Open(filepath); err == nil {
			defer imgFile.Close()
			if config, _, err := image.DecodeConfig(imgFile); err == nil {
				metadata.Width = config.Width
				metadata.Height = config.Height
			}
		}
	}

	// GPS coordinates - returns (float64, float64, bool)
	if lat, lon, ok := m.GPS(); ok {
		metadata.GPSLatitude = lat
		metadata.GPSLongitude = lon
		metadata.HasGPS = true
	}

	return metadata
}
