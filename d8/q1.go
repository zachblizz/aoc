package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

)

// color distinction
const (
	black       = 0
	white       = 1
	transparent = 2
)

func getDim(dim string) (int, int) {
	s := strings.Split(dim, "x")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])

	return x, y
}

func createLayers(input, dim string) ([][]int, int, int) {
	x, y := getDim(dim)
	layers := [][]int{}
	l := 0

	for _, rawByte := range input {
		b, _ := strconv.Atoi(string(rawByte))

		if len(layers) == 0 {
			layers = [][]int{[]int{b}}
		} else if len(layers[l]) < x {
			layers[l] = append(layers[l], b)
		} else {
			l++
			layers = append(layers, []int{b})
		}
	}

	return layers, x, y
}

func printLayers(layers [][]int, tall int) {
	for i, row := range layers {
		if i%tall == 0 {
			fmt.Printf("layer %v:\n", i)
		}
		// fmt.Println(i, row)

		for _, col := range row {
			fmt.Print(col)
		}

		fmt.Println()
	}
}

func getSmallestLayer(layers [][]int, tall int) int {
	zMap := make(map[int]int)
	layer := 0

	zeros := math.MaxInt64
	smallLayer := 0

	for i, row := range layers {
		if i%tall == 0 {
			layer++
		}

		for _, b := range row {
			if b == 0 {
				zMap[layer]++
			}
		}
	}

	for l, zs := range zMap {
		if zeros > zs {
			zeros = zs
			smallLayer = l
		}
	}

	return smallLayer
}

func calculateOnesAndToos(layers [][]int, tall, smallestLayer int) int {
	layer := 0
	histo := make(map[int]int)

	for i, row := range layers {
		if i%tall == 0 {
			layer++
		}

		if layer == smallestLayer {
			for _, b := range row {
				if b == 2 || b == 1 {
					histo[b]++
				}
			}
		} else if layer > smallestLayer {
			break
		}
	}

	fmt.Println(histo)

	return histo[1] * histo[2]
}

func drawPixels(layers [][]int, x, y int) [][]int {
	pixels := make([][]int, y)

	// initialize all pixel rows with transparent
	for i := 0; i < y; i++ {
		pixels[i] = []int{}

		for j := 0; j < x; j++ {
			pixels[i] = append(pixels[i], transparent)
		}
	}

	layer := 0
	for i, pRow := range layers {
		layer = i % y
		for j, pc := range pRow {
			if pixels[layer][j] == transparent {
				pixels[layer][j] = pc
			}
		}
	}

	return pixels
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	dim := "25x6"

	// input := "0222112222120000"
	// dim = "2x2"

	// 25x6 image
	// layer consists of 6 rows with 25 numbers in them
	layers, x, y := createLayers(string(input), dim)
	// printLayers(layers, y)
	// smallestLayer := getSmallestLayer(layers, y)
	// fmt.Println(calculateOnesAndToos(layers, y, smallestLayer))

	printLayers(drawPixels(layers, x, y), y)
}
