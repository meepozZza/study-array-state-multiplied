package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter array via comma: ")
	arrayString := readStringInput()
	array := strings.Split(arrayString, ",")
	intArray := make([]int, len(array))

	for i, v := range array {
		v, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		intArray[i] = v
	}

	fmt.Print("Enter itterations count: ")
	itterations := readIntInput()

	fmt.Print("Enter multiplier: ")
	multiplier := readIntInput()

	finalState := getFinalState(intArray, itterations, multiplier)

	fmt.Print("Result: ", finalState, "\n")
}

func readStringInput() (val string) {
	_, err := fmt.Scanf("%s", &val)

	if err != nil {
		panic(err)
	}

	return
}

func readIntInput() (val int) {
	_, err := fmt.Scanf("%d", &val)

	if err != nil {
		panic(err)
	}

	return
}

func getFinalState(nums []int, k int, multiplier int) []int {
	minValueIndex := 0

	for k > 0 {
		for index, value := range nums {
			if value < nums[minValueIndex] {
				minValueIndex = index
			}
		}
		nums[minValueIndex] *= multiplier
		minValueIndex = 0
		k--
	}

	return nums
}
