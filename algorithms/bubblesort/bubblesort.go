package main

import "fmt"

// only good for almost sorted lists

func main() {
	numbers := []int{4, 5, 3, 0, 1, 2}
	fmt.Println("List of numbers:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sort: ", numbers)
}

func bubbleSort(numbers []int) {
	N := len(numbers)
	for i := 0; i < N; i++ {
		fmt.Println("Doing a sweep:", numbers)
		if !sweep(numbers, i) {
			return
		}
	}
}

// prevPasses is used to optimize the loop. After every sweep, the largest number is already at the end of slice
// so we can skip the last n numbers after every n sweeps as they are already in their final places.

// didSwap is returned to let the calling loop know when there was no swap, indicating the slice is sorted.
func sweep(numbers []int, prevPasses int) bool {
	N := len(numbers)
	firstIndex := 0
	secondIndex := 1
	didSwap := false

	for secondIndex < (N - prevPasses) {
		firstNumber := numbers[firstIndex]
		secondNumber := numbers[secondIndex]
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
