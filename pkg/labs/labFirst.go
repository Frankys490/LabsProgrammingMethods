package labs

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Lab1A() {
	inputData := strings.Split(scanFile("src/Lab1A.txt"), ";")
	inputData = inputData[1 : len(inputData)-1]
	var queue []string

	for _, elem := range inputData {
		signAndNumber := strings.Split(elem, " ")
		switch signAndNumber[0] {
		case "-":
			fmt.Println(queue[0])
			queue = queue[1:]
		case "+":
			queue = append(queue, signAndNumber[1])
		case "*":
			if len(queue)%2 == 0 {
				index := len(queue) / 2
				queue = append(queue[:index+1], queue[index:]...)
				queue[index] = signAndNumber[1]
			} else {
				index := len(queue)/2 + 1
				queue = append(queue[:index+1], queue[index:]...)
				queue[index] = signAndNumber[1]
			}
		}
	}
}

func Lab1D() {
	inputData := strings.Split(strings.ReplaceAll(scanFile("src/Lab1D.txt"), ";", ""), " ")
	n, err := strconv.Atoi(inputData[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(inputData[1])
	if err != nil {
		panic(err)
	}
	c := 0

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	for i := 0; i < m+n-1; i++ {
		for j := 0; j < n; j++ {
			if ((i - j) > -1) && ((i - j) < m) {
				matrix[j][i-j] += c
				c += 1
			}
		}
	}

	for idx, _ := range matrix {
		fmt.Println(matrix[idx])
	}
}

func Lab1I() {
	inputData := strings.Split(scanFile("src/Lab1I.txt"), ";")
	inputData = inputData[:len(inputData)-1]
	studentsMap := make(map[int][]string)
	var keys []int

	for _, elem := range inputData {
		numberAndLastName := strings.Split(elem, " ")
		number, err := strconv.Atoi(numberAndLastName[0])
		if err != nil {
			panic(err)
		}
		studentsMap[number] = append(studentsMap[number], numberAndLastName[1])
	}
	for key, _ := range studentsMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		for _, elem := range studentsMap[key] {
			fmt.Println(key, elem)
		}
	}
}
