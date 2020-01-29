package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

)

func getDim(dim string) (int, int) {
	s := strings.Split(dim, "x")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])

	return x, y
}

func createLayers(input, dim string) ([][]int, int) {
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

	return layers, y
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

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	dim := "25x6"

	// 25x6 image
	// layer consists of 6 rows with 25 numbers in them
	layers, tall := createLayers(string(input), dim)
	printLayers(layers, tall)
	// smallestLayer := getSmallestLayer(layers, tall)
	// fmt.Println(calculateOnesAndToos(layers, tall, smallestLayer))

}
