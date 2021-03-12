package main

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "image/png"

	"github.com/disintegration/imaging"
)

func checkAndRotate(imagePath string) {
	width, height := getImageDimension(imagePath)
	if width > height {
		rotate270(imagePath)
	}
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)

	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}

func rotate270(imagePath string) {

	src, err := imaging.Open(imagePath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	dst := imaging.Rotate270(src)

	// Save the resulting image as JPEG.
	err = imaging.Save(dst, imagePath)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
