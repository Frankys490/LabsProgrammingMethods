package labs

import (
	"fmt"
	"strconv"
	"strings"
)

func Lab4A() {
	counter := 1
	alphabet := make(map[string]int)
	alphabet["a"] = 0
	alphabet["b"] = 1
	alphabet["c"] = 2
	alphabet["d"] = 3
	alphabet["e"] = 4
	alphabet["f"] = 5
	alphabet["g"] = 6
	alphabet["h"] = 7

	inputData := strings.Split(strings.Trim(scanFile("src/Lab4A.txt"), ";"), " ")
	firstCoordinatesString := strings.Split(inputData[0], "")
	secondCoordinatesString := strings.Split(inputData[1], "")
	firstNumberToConvert, err := strconv.Atoi(firstCoordinatesString[1])
	if err != nil {
		panic(err)
	}
	secondNumberToConvert, err := strconv.Atoi(secondCoordinatesString[1])
	if err != nil {
		panic(err)
	}
	horses := [][]int{{firstNumberToConvert - 1, alphabet[firstCoordinatesString[0]]},
		{secondNumberToConvert - 1, alphabet[secondCoordinatesString[0]]}}

	firstPoints := nextMove(horses[0])
	secondPoints := nextMove(horses[1])
	if (((horses[0][0] - horses[1][0]) + (horses[0][1] - horses[1][1])) % 2) != 0 {
		fmt.Println(-1)
	} else if comparison(firstPoints, secondPoints) {
		fmt.Println(counter)
	} else {
		for {
			counter++
			firstLenToDel := len(firstPoints)
			secondLenToDel := len(secondPoints)
			for _, elem := range firstPoints {
				firstPoints = append(firstPoints, nextMove(elem)...)
			}
			for _, elem := range secondPoints {
				secondPoints = append(secondPoints, nextMove(elem)...)
			}
			firstPoints = firstPoints[firstLenToDel:]
			secondPoints = secondPoints[secondLenToDel:]
			if comparison(firstPoints, secondPoints) {
				fmt.Println(counter)
				break
			}
		}
	}
}

func nextMove(edge []int) [][]int {
	var points [][]int
	if edge[0]-2 >= 0 {
		if edge[1]-1 >= 0 {
			point := []int{edge[0] - 2, edge[1] - 1}
			points = append(points, point)
		}
		if edge[1]+1 < 8 {
			point := []int{edge[0] - 2, edge[1] + 1}
			points = append(points, point)
		}
	}
	if edge[0]+2 < 8 {
		if edge[1]-1 >= 0 {
			point := []int{edge[0] + 2, edge[1] - 1}
			points = append(points, point)
		}
		if edge[1]+1 < 8 {
			point := []int{edge[0] + 2, edge[1] + 1}
			points = append(points, point)
		}
	}
	if edge[1]-2 >= 0 {
		if edge[0]-1 >= 0 {
			point := []int{edge[0] - 1, edge[1] - 2}
			points = append(points, point)
		}
		if edge[0]+1 < 8 {
			point := []int{edge[0] + 1, edge[1] - 2}
			points = append(points, point)
		}
	}
	if edge[1]+2 < 8 {
		if edge[0]-1 >= 0 {
			point := []int{edge[0] - 1, edge[1] + 2}
			points = append(points, point)
		}
		if edge[0]+1 < 8 {
			point := []int{edge[0] + 1, edge[1] + 2}
			points = append(points, point)
		}
	}
	return points
}

func comparison(firstPoints [][]int, secondPoints [][]int) bool {
	equal := false
	for _, elemOfFirstPoints := range firstPoints {
		for _, elemOfSecondPoints := range secondPoints {
			if elemOfFirstPoints[0] == elemOfSecondPoints[0] && elemOfFirstPoints[1] == elemOfSecondPoints[1] {
				equal = true
			}
		}
	}
	return equal
}
