package labs

import (
	"fmt"
	"strconv"
	"strings"
)

func Lab3A() {
	inputData := strings.Split(scanFile("src/Lab3A.txt"), ";")
	inputData = inputData[:len(inputData)-1]
	lenOfFirstSequence, err := strconv.Atoi(inputData[0])
	if err != nil {
		panic(err)
	}
	firstSequence := strings.Split(inputData[1], " ")
	lenOfSecondSequence, err := strconv.Atoi(inputData[2])
	if err != nil {
		panic(err)
	}
	secondSequence := strings.Split(inputData[3], " ")
	if lenOfFirstSequence != len(firstSequence) || lenOfSecondSequence != len(secondSequence) {
		panic("Проверьте правльность введенных данных")
	}

	chart := make([][]int, lenOfFirstSequence)
	for i := 0; i < lenOfFirstSequence; i++ {
		for j := 0; j < lenOfSecondSequence; j++ {
			chart[i] = append(chart[i], 0)
		}
	}

	for idxFirst, elemFirst := range firstSequence {
		for idxSecond, elemSecond := range secondSequence {
			if elemFirst == elemSecond {
				if idxFirst == 0 || idxSecond == 0 {
					chart[idxFirst][idxSecond] = 1
				} else {
					chart[idxFirst][idxSecond] = chart[idxFirst-1][idxSecond-1] + 1
				}
			} else {
				chart[idxFirst][idxSecond] = 0
			}
		}
	}
	max := 0
	maxIdx := 0
	var result []string
	for idxFirst, _ := range chart {
		for _, elem := range chart[idxFirst] {
			if elem > max {
				max = elem
				maxIdx = idxFirst
			}
		}
	}
	for i := max; i > 0; i-- {
		result = append(result, firstSequence[maxIdx])
		maxIdx--
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i], " ")
	}
}

func Lab3D() {
	inputData := strings.Split(scanFile("src/Lab3D.txt"), ";")
	inputData = inputData[:len(inputData)-1]
	lengthOfSequence, err := strconv.Atoi(inputData[0])
	if err != nil {
		panic(err)
	}
	stringSequence := strings.Split(inputData[1], " ")
	if len(stringSequence) != lengthOfSequence {
		panic("Wrong data!")
	}
	var intSequence []int

	for _, elem := range stringSequence {
		number, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		intSequence = append(intSequence, number)
	}

	var subsequence []int
	for i := 0; i < lengthOfSequence; i++ {
		subsequence = append(subsequence, 0)
	}

	for i := 0; i < lengthOfSequence; i++ {
		for j := 0; j < i; j++ {
			if intSequence[j] < intSequence[i] && subsequence[i] < subsequence[j] {
				subsequence[i] = subsequence[j]
			}
		}
		subsequence[i] += 1
	}

	var final []int
	maxLengthOfSubsequence := subsequence[len(subsequence)-1]
	for i := len(subsequence) - 1; i > 0; i-- {
		if subsequence[i] == subsequence[i-1] && subsequence[i] == maxLengthOfSubsequence {
			final = append(final, intSequence[i])
			maxLengthOfSubsequence--
		}
		if subsequence[i] != subsequence[i-1] && i == 1 {
			final = append(final, intSequence[i-1])
		}
	}
	for i := len(final) - 1; i >= 0; i-- {
		fmt.Print(final[i], " ")
	}
	fmt.Println()
}

func Lab3I() {
	inputData := strings.Split(scanFile("src/Lab3I.txt"), ";")
	inputData = inputData[:len(inputData)-1]
	possibleWeight, err := strconv.Atoi((strings.Split(inputData[0], " "))[0])
	if err != nil {
		panic(err)
	}
	numberOfBars, err := strconv.Atoi((strings.Split(inputData[0], " "))[1])
	if err != nil {
		panic(err)
	}
	stringWeights := strings.Split(inputData[1], " ")
	var intWeights []int

	if numberOfBars != len(stringWeights) {
		panic("Wrong data!")
	}

	for _, elem := range stringWeights {
		weight, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		intWeights = append(intWeights, weight)
	}

	chart := make([][]int, numberOfBars+1)
	for i := 0; i < numberOfBars+1; i++ {
		for j := 0; j < possibleWeight+1; j++ {
			chart[i] = append(chart[i], 0)
		}
	}

	for i := 0; i <= numberOfBars; i++ {
		for j := 0; j <= possibleWeight; j++ {
			if i == 0 || j == 0 {
				chart[i][j] = 0
			} else {
				if intWeights[i-1] > j {
					chart[i][j] = chart[i-1][j]
				} else {
					previousMax := chart[i-1][j]
					formula := intWeights[i-1] + chart[i-1][j-intWeights[i-1]]
					if previousMax > formula {
						chart[i][j] = previousMax
					} else {
						chart[i][j] = formula
					}
				}
			}
		}
	}
	fmt.Println(chart[numberOfBars][possibleWeight])
}
