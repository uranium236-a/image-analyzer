package cli

import (
	"fmt"
	analyzemetadata "image-analyzer/analyze-metadata"
	"strings"
)

func Start() {
	for {
		command, err := ReadInput("\n>> ")

		if err != nil {
			fmt.Println(err)
			continue
		}

		subcommands := strings.Fields(command)

		if len(subcommands) == 0 {
			continue
		}

		switch subcommands[0] {
		case "help":
			help()
		case "rgb":
			if len(subcommands) < 2 {
				fmt.Println("Error: image path not found")
			} else {
				processImage(subcommands[1])
			}

		case "metadata":
			if len(subcommands) < 2 {
				fmt.Println("Error: image path not found")
			} else {
				metaData(subcommands[1])
			}

		case "clear":
			clearScreen()

		case "exit":
			return

		default:
			fmt.Println("Invalid command")
		}
	}
}

func help() {
	fmt.Println("\nrgb <filepath> :- shows RGB info")
	fmt.Println("metadata <filepath> :- metadata of the image")
	fmt.Println("clear :- clear screen")
	fmt.Println("exit :- exit the program")
}

func metaData(filepath string) {
	metadata := analyzemetadata.Analyze(filepath)

	if metadata.Error != nil {
		fmt.Println("Error: ", metadata.Error)
	} else {
		fmt.Println("\nCamera Make: ", metadata.CameraMake)
		fmt.Println("Camera Model: ", metadata.CameraModel)
		fmt.Println("Date/Time: ", metadata.DateTime)
		fmt.Println("Focal Length: ", metadata.FocalLength)
		fmt.Println("ISO: ", metadata.ISO)
		fmt.Println("Width: ", metadata.Width)
		fmt.Println("Height: ", metadata.Height)
		fmt.Println("Latitude: ", metadata.GPSLatitude)
		fmt.Println("Longitude: ", metadata.GPSLongitude)
		fmt.Println("GPS: ", metadata.HasGPS)
	}
}
