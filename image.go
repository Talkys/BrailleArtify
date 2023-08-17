package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func loadImage(filePath string, threshold int) ([][]uint8, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	matrix := make([][]uint8, height)
	for y := 0; y < height; y++ {
		matrix[y] = make([]uint8, width)
		for x := 0; x < width; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			grayValue := uint8((r + r + r) / 3)
			if grayValue > uint8(threshold) {
				matrix[y][x] = 1
			} else {
				matrix[y][x] = 0
			}
		}
	}

	return matrix, nil
}

func extendMatrix(matrix [][]uint8) [][]uint8 {
	rows := len(matrix)
	cols := len(matrix[0])

	extraRows := (4 - rows%4) % 4
	extraCols := (2 - cols%2) % 2

	nRows, nCols := rows+extraRows, cols+extraCols

	nMatrix := make([][]uint8, nRows)
	for i := range nMatrix {
		nMatrix[i] = make([]uint8, nCols)
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			nMatrix[y][x] = matrix[y][x]
		}
	}

	return nMatrix
}

func ImgToMat(imagePath string, threshold int) [][]uint8 {
	matrix, err := loadImage(imagePath, threshold)
	if err != nil {
		return nil
	}
	return extendMatrix(matrix)
}
