package main

import (
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitTiles2(matrix [][]uint8, Xs int, Ys int) [][][]uint8 {
	rows := len(matrix)
	cols := len(matrix[0])

	var tiles [][][]uint8

	for i := 0; i < rows; i += Ys {
		for j := 0; j < cols; j += Xs {
			tileRows := min(Ys, rows-i)
			tileCols := min(Xs, cols-j)

			tile := make([][]uint8, 4) // 4 linhas
			for y := 0; y < 4; y++ {
				tile[y] = make([]uint8, 2) // 2 colunas
				for x := 0; x < 2; x++ {
					if y < tileRows && x < tileCols {
						tile[y][x] = matrix[i+y][j+x]
					} else {
						tile[y][x] = 0 // Preencha com algum valor padrão ou lógica adequada
					}
				}
			}

			tiles = append(tiles, tile)
		}
	}

	return tiles
}

/*
func splitTiles(matrix [][]uint8, Xs int, Ys int) [][][]uint8 {
	rows := len(matrix)
	cols := len(matrix[0])

	tiles := make([][][]uint8, 0)

	for i := 0; i < rows; i += Ys {
		for j := 0; j < cols; j += Xs {
			var tile [][]uint8
			for y := i; y < i+Ys && y < rows; y++ {
				row := make([]uint8, 0)
				for x := j; x < j+Xs && x < cols; x++ {
					row = append(row, matrix[y][x])
				}
				tile = append(tile, row)
			}
			tiles = append(tiles, tile)
		}
	}

	return tiles
}
*/

func mapChar(tile [][]uint8) rune {
	binary := ""
	binary += discretize(tile[3][1])
	binary += discretize(tile[3][0])
	binary += discretize(tile[2][1])
	binary += discretize(tile[1][1])
	binary += discretize(tile[0][1])
	binary += discretize(tile[2][0])
	binary += discretize(tile[1][0])
	binary += discretize(tile[0][0])

	base := 0x2800
	decimalValue, _ := strconv.ParseInt(binary, 2, 0)
	return rune(base + int(decimalValue))
}

//Jeito mais fácil de converter só 1 e 0
func discretize(v uint8) string {
	if v == 1 {
		return "1"
	} else {
		return "0"
	}
}

func generate(filepath string, threshold int) string {

	frame := ""
	if !(strings.HasSuffix(filepath, ".png") ||
		strings.HasSuffix(filepath, ".jpg") ||
		strings.HasSuffix(filepath, ".jpeg")) {
		return ""
	}
	matrix := ImgToMat(filepath, threshold)
	X := 2
	Y := 4

	//startTime := time.Now()
	tiles := splitTiles2(matrix, X, Y)
	//elapsedTime := time.Since(startTime)
	//fmt.Println(elapsedTime)
	var letras []rune

	for _, tile := range tiles {
		letras = append(letras, mapChar(tile))
	}

	cols := len(matrix[0]) / X
	chars := len(letras)
	for i := 0; i < chars; i += cols {
		end := i + cols
		if end > chars {
			end = chars
		}

		frame += (string(letras[i:end]) + "\n")
	}

	return frame
}
