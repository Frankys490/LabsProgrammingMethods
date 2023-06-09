package labs

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Lab2A() {
	inputData := strings.Split(scanFile("src/Lab2A.txt"), "")
	inputData = inputData[:len(inputData)-1]
	var arrayOfLetters []string
	var arrayOfNumbers []string
	for _, elem := range inputData {
		r := []rune(elem)
		if unicode.IsDigit(r[0]) {
			arrayOfNumbers = append(arrayOfNumbers, elem)
		} else {
			arrayOfLetters = append(arrayOfLetters, elem)
		}
	}

	arrayOfLettersCombinations := removeDuplicates(combinations(arrayOfLetters))
	arrayOfNumbersCombinations := removeDuplicates(combinations(arrayOfNumbers))

	var arrayOfAllPossibleCombinations []string
	for _, elemLetter := range arrayOfLettersCombinations {
		for _, elemNumber := range arrayOfNumbersCombinations {
			r := []rune(elemLetter)
			arrayOfAllPossibleCombinations = append(arrayOfAllPossibleCombinations, string(r[0])+elemNumber+string(r[1])+string(r[1]))
		}
	}
	fmt.Println(len(arrayOfAllPossibleCombinations))
	for _, elem := range arrayOfAllPossibleCombinations {
		fmt.Println(elem)
	}
}

func combinations(arrayStart []string) []string {
	var arrayFinish []string
	for idx, _ := range arrayStart {
		var arrayForComb []string
		arrayForComb = append(arrayForComb, arrayStart[:idx]...)
		arrayForComb = append(arrayForComb, arrayStart[idx+1:]...)
		for i := 0; i < 2; i++ {
			arrayFinish = append(arrayFinish, arrayStart[idx]+arrayForComb[0]+arrayForComb[1])
			arrayForComb[0], arrayForComb[1] = arrayForComb[1], arrayForComb[0]
		}
	}
	return arrayFinish
}

func removeDuplicates(array []string) []string {
	keys := make(map[string]int)
	var finishArray []string
	for _, elem := range array {
		keys[elem] += 1
	}
	for key, _ := range keys {
		finishArray = append(finishArray, key)
	}
	return finishArray
}
func Lab2H() {
	inputData := strings.Split(scanFile("src/Lab2H.txt"), ";")
	numberOfSegments, err := strconv.Atoi(strings.Split(inputData[0], " ")[1])
	if err != nil {
		panic(err)
	}
	inputData = inputData[1 : len(inputData)-1]
	var lengthOfWires []int
	for _, elem := range inputData {
		length, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		lengthOfWires = append(lengthOfWires, length)
	}
	sort.Ints(lengthOfWires)

	sumOfLengths := 0
	for _, elem := range lengthOfWires {
		sumOfLengths += elem
	}
	if sumOfLengths < numberOfSegments {
		fmt.Println(0)
		return
	}

	min, max := 0, lengthOfWires[len(lengthOfWires)-1]
	var mid int
	for max-min > 1 {
		midNumberOfSegments := 0
		mid = (max + min) / 2
		for _, elem := range lengthOfWires {
			midNumberOfSegments += elem / mid
		}
		if midNumberOfSegments < numberOfSegments {
			max = mid - 2
		} else {
			min = mid
		}
	}
	fmt.Println(mid)
}

func Lab2I() {
	inputData := strings.Split(strings.ReplaceAll(scanFile("src/Lab2I.txt"), ";", ""), "")
	var arrayOfBlocks [][]string
	for len(inputData) > 10 {
		min, max := len(inputData), 0
		arrayToCheck := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		for idxInputData := len(inputData) - 1; idxInputData >= 0; idxInputData-- {
			for idx, elem := range arrayToCheck {
				if inputData[idxInputData] == elem {
					if max < idxInputData {
						max = idxInputData
					}
					if min > idxInputData {
						min = idxInputData
					}
					arrayToCheck = append(arrayToCheck[:idx], arrayToCheck[idx+1:]...)
				}
			}

		}
		arrayOfBlocks = append(arrayOfBlocks, inputData[min:max+1])
		inputData = inputData[:min]
	}
	arrayOfBlocks = append(arrayOfBlocks, inputData)
	var arrayFinish []string
	var numberToDel string
	for idxArrayOfBlocks := len(arrayOfBlocks) - 1; idxArrayOfBlocks >= 0; idxArrayOfBlocks-- {
		if numberToDel != "" {
			for idx, elem := range arrayOfBlocks[idxArrayOfBlocks] {
				if numberToDel == elem {
					arrayOfBlocks[idxArrayOfBlocks] = arrayOfBlocks[idxArrayOfBlocks][idx+1:]
				}
			}
		}
		stringToCheck := strings.Join(arrayOfBlocks[idxArrayOfBlocks], "")
		arrayToCheck := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		for _, elem := range arrayToCheck {
			if idxArrayOfBlocks == len(arrayOfBlocks)-1 && elem == "0" {
				continue
			}
			if !strings.Contains(stringToCheck, elem) {
				arrayFinish = append(arrayFinish, elem)
				numberToDel = elem
				break
			}
		}
	}
	fmt.Println(strings.Join(arrayFinish, ""))
}
