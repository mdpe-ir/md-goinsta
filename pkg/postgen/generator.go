package postgen

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
)

type PostGenConfig struct {
	InstagramUsername string
	Content           string
	SlidesCount       int
	Index             int
}

func PostGen(config PostGenConfig) error {

	// Load template image
	templateImage, err := GetImageFromFilePath("assets/templates/temp.png")
	if err != nil {
		return err
	}

	// Create a new image with the size of the original image
	templateImageDst := image.NewNRGBA(templateImage.Bounds())

	// Draw the base image onto the new image
	draw.Draw(templateImageDst, templateImage.Bounds(), templateImage, image.Point{0, 0}, draw.Over)

	// Add username to text
	usernameWriter, usernameColor, err := GenerateTextWriter(config.InstagramUsername, "assets/fonts/Vazirmatn-Medium.ttf", "#000000", 32) // Get Color And Size From Figma
	if err != nil {
		return err
	}
	defer usernameWriter.Close()
	usernameWriter.Write(templateImageDst, image.Point{X: 168, Y: 123}, usernameColor) // Get Position From Figma

	// Save the output image
	outputFile, err := os.Create("output_image.png") // TODO: Change This Name Later
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, templateImageDst)
	if err != nil {
		return err
	}

	return nil

}
