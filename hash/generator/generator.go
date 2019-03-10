package generator

import "github.com/hermes7308/golobal/symmetric"

const (
	HASH_SIZE     = 55
	RESIZE_WIDTH  = 150
	RESIZE_HEIGHT = 200
	NGRAYBLOCK    = 85
)

var MEAN = []float32{
	149.1, 145.6,
	144.7, 144.1, 144.2, 145.1,
	148.3, 145.2, 138.3, 136.5,
	137.4, 136.4, 137.6, 143.9,
	144.0, 136.6, 136.5, 138.7,
	137.0, 136.2, 142.7, 142.5,
	135.4, 136.2, 138.6, 136.9,
	135.2, 141.2, 140.5, 133.4,
	133.6, 135.2, 134.0, 133.2,
	139.4, 138.5, 132.0, 131.6,
	132.6, 131.9, 131.7, 137.5,
	148.1, 143.9, 143.4, 143.8,
	143.2, 142.8, 146.4, 148.1,
	144.5, 143.2, 143.2, 144.0,
	147.5, 144.1, 137.1, 136.7,
	137.1, 136.8, 143.0, 142.5,
	135.9, 137.5, 138.4, 136.0,
	141.2, 140.5, 134.1, 135.6,
	136.4, 134.3, 139.5, 138.2,
	132.0, 132.3, 132.8, 132.0,
	137.2, 146.3, 142.2, 141.9,
	142.2, 141.4, 144.8}

func GetGrayBlock(imageBytes []uint8, width int, height int) []float32 {
	var numBlock int
	var xaxisSize int
	var yaxisSize int
	var xaxisIndex int
	var yaxisIndex int
	var indexHist int
	histogram := make([]float32, NGRAYBLOCK)
	histogram2 := make([]float32, NGRAYBLOCK)
	pixelN := make([]float32, NGRAYBLOCK)
	pixelCnt := width * height
	var grayResult float32

	for i := 0; i < 2; i++ {
		if i == 0 {
			numBlock = 7
		} else {
			numBlock = 6
		}

		xaxisSize = width / numBlock
		yaxisSize = height / numBlock
		if xaxisSize == 0 {
			xaxisSize = 1
		}
		if yaxisSize == 0 {
			yaxisSize = 1
		}

		for j := 0; j < NGRAYBLOCK; j++ {
			histogram[j] = 0
			pixelN[j] = 0
		}

		for y := 0; y < pixelCnt; y += 3 {
			// grayResult = (float) (red[y] + green[y] + blue[y]) / (float) 3.0;
			blue := imageBytes[y+0]
			green := imageBytes[y+1]
			red := imageBytes[y+2]

			grayResult = float32(int((blue + green + red) / 3))
			xaxisIndex = y % width
			yaxisIndex = y / width
			xaxisIndex /= xaxisSize
			yaxisIndex /= yaxisSize
			if yaxisIndex > numBlock-1 {
				yaxisIndex = numBlock - 1
			}
			if xaxisIndex > numBlock-1 {
				xaxisIndex = numBlock - 1
			}

			indexHist = xaxisIndex + yaxisIndex*numBlock

			// exception
			if (indexHist < 0) || (indexHist > numBlock*numBlock-1) {
				//TODO error log
			}

			histogram[indexHist] += grayResult
			pixelN[indexHist]++
		}

		numBlock2 := numBlock * numBlock
		for j := 0; j < numBlock2; j++ {
			if pixelN[j] != 0.0 {
				histogram2[j+49*i] = histogram[j] / pixelN[j]
			}
		}
	}

	return histogram2
}

func GetGlobalHash(grayBlock []float32, hashSize int) int64 {
	// calc
	interm := make([]float32, NGRAYBLOCK)
	result := make([]float32, NGRAYBLOCK)

	for i := 0; i < NGRAYBLOCK; i++ {
		interm[i] = grayBlock[i] - MEAN[i]
	}

	for i := 0; i < 55; i++ {
		result[i] = 0.0
		for j := 0; j < NGRAYBLOCK; j++ {
			result[i] += interm[j] * symmetric.METRIC[j][i]
		}
	}

	return CalculateHashValue(result, hashSize)
}

func CalculateHashValue(result []float32, hashSize int) int64 {
	var hashValue int64 = 0
	var tempHash int64 = 0

	// hashValue initialize
	hashValue = hashValue & 0x0000000000000000

	for i := 0; i < hashSize; i++ {
		if result[i] > 0.0 {
			tempHash = 0x0000000000000001 << uint64(hashSize-1-i)
		} else {
			tempHash = 0x0000000000000000
		}

		hashValue = hashValue | tempHash
	}

	return hashValue
}
