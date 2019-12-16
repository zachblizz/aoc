package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x int
	y int
}

func constructPoints(wire []string) (map[string]point, []string) {
	wireMap := make(map[string]point)
	order := []string{}
	xPos := 0
	yPos := 0

	for _, dirDist := range wire {
		dir := string(dirDist[0])
		dist, _ := strconv.Atoi(dirDist[1:len(dirDist)])

		switch dir {
		case "U":
			yPos += dist
			break
		case "D":
			yPos -= dist
			break
		case "L":
			xPos -= dist
			break
		case "R":
			xPos += dist
			break
		default:
			fmt.Printf("Dir %v not found\n", dir)
			return nil, nil
		}

		key := strconv.Itoa(xPos) + strconv.Itoa(yPos)
		wireMap[key] = point{xPos, yPos}
		order = append(order, key)
	}

	return wireMap, order
}

func getIntersections(w1P, w2P map[string]point) point {
	return point{}
}

func main() {
	w1 := []string{"R8", "U5", "L5", "D3"}
	w2 := []string{"U7", "R6", "D4", "L4"}
	w1Points, w1Order := constructPoints(w1)
	fmt.Println(w1Points, w1Order)
	w2Points, w2Order := constructPoints(w2)
	fmt.Println(w2Points, w2Order)
	// getIntersections(w1Points, w2Points)
}
