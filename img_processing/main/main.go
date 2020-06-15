package main

import (
	"fmt"
	img "image"
	imgPNG "image/png"
	"os"
)

func openImg(filePath string) img.Image {

	fileReader, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("\nImage file cannot be opened\n")
		panic(err)

	} else {
		fmt.Printf("\nImage file was opened successfully\n")

		// Close file reader
		defer fileReader.Close()

		// Decoding image file
		imageFile, err := imgPNG.Decode(fileReader)

		if err != nil {
			fmt.Printf("\nCannot incode image file\n")
			panic(err)
		} else {

			// Return image file
			return imageFile
		}
	}
}

func saveFile(outFilePath string, imageFile img.Image) {

	outFile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Printf("\nCannot create image file\n")
		panic(err)
	} else {
		// A defer statement defers the execution of a function until the surrounding function returns.
		// The deferred call's arguments are evaluated immediately, but the function call is not executed until the
		// surrounding function returns.
		defer outFile.Close()

		imgPNG.Encode(outFile, imageFile)

		fmt.Printf("\nImage file was successfully created\n\n")
	}

}

func splitImage(imageFile img.Image) img.Image {

	print("\nSplitting up image\n")

	newImgW := 25 // 25 px
	newImgH := 25 // 25 px

	// Initialize empty image file
	newImage := img.NewRGBA(img.Rect(0, 0, newImgW, newImgH))

	for x := 0; x < newImgW; x++ {
		for y := 0; y < newImgH; y++ {

			// Get original pixel color
			tmpColor := imageFile.At(x, y)

			// Set pixel color in the new image file
			newImage.Set(x, y, tmpColor)
		}
	}

	return newImage
}

func main() {
	filePath := "../images/numbers.png"
	outFilePath := "../images/new.png"

	// Declare variable to hold Image object
	var imageFile img.Image
	imageFile = openImg(filePath)

	if imageFile != nil {

		// TEST
		// ========================================================================= //
		imageBounds := imageFile.Bounds()
		imageWidth := imageBounds.Max.X
		imageHight := imageBounds.Max.Y
		fmt.Printf("Image Width: %d --- Hight: %d\n\n", imageWidth, imageHight)
		// ========================================================================= //

		// Split image
		newImg := splitImage(imageFile)

		if newImg != nil {
			// Save image file
			saveFile(outFilePath, newImg)
		}
	}

	// http://tech.nitoyon.com/en/blog/2015/12/31/go-image-gen/
	// https://stackoverflow.com/questions/45226991/golang-image-colormodel
	//https://stackoverflow.com/questions/8697095/how-to-read-a-png-file-in-color-and-output-as-gray-scale-using-the-go-programmin#

	// https://golang.org/pkg/image/
	// 
}

