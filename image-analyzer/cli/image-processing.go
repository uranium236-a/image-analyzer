package cli

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func processImage(imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	r, g, b, a := img.At(100, 100).RGBA()
	fmt.Printf("Pixel: R=%d G=%d B=%d A=%d\n", r, g, b, a)
}
