package main

import (
	"image/gif"
	"log"
	"os"
)

func main() {
	gifFilePath := "./resources/image/animated1.gif"
	//pngFilePath := "./resources/image/Lena.png"

	// image
	imageFile, err := os.Open(gifFilePath)
	if err != nil {
		log.Panic("Can not read gif image ", imageFile, err)
		return
	}
	defer imageFile.Close()

	srcImage, err := gif.Decode(imageFile)
	if err != nil {
		log.Panic("Can not decode", err)
		return
	}

	log.Println(gifFilePath, srcImage.At(0, 0))

}
