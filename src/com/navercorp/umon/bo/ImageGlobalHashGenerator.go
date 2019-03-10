package imageGlobalHashGenerator

import (
	"errors"
	"gocv.io/x/gocv"
	"log"
)

func GetGlobalHash(originalMat gocv.Mat, hashSize int, resizeWidth int, resizeHeight int) {

}

func GetGrayImage(filePath string) (gocv.Mat, error) {
	originalMat := gocv.IMRead(filePath, gocv.IMReadGrayScale)
	if originalMat.Empty() {
		log.Panic("Can not read image:", filePath)
		return originalMat, errors.New("Can not read Image:" + filePath)
	}

	return originalMat, nil
}
