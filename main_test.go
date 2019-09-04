package main

import (
	"testing"
)

func Test_bubble(t *testing.T) {
	var testArray = []int{435, 543, 435, 342, 543, 34, 532, 56594, 26, 73864, 267, 2, 1, 77, 8, 368, 686, 864, 3, 4}
	BubleSort(&testArray)

	for i := 0; i < len(testArray)-1; i++ {
		if testArray[i] > testArray[i+1] {
			t.Error("Not sorted!")
		}
	}
	//fmt.Println("Bubble sort: ", testArray)
}

func Test_Quick(t *testing.T) {
	var testArray = []int{435, 543, 435, 342, 543, 34, 532, 56594, 26, 73864, 267, 2, 1, 77, 8, 368, 686, 864, 3, 4}
	QuickSort(&testArray, 0, len(testArray)-1)

	for i := 0; i < len(testArray)-1; i++ {
		if testArray[i] > testArray[i+1] {
			t.Error("Not sorted!")
		}
	}
	//fmt.Println("Quick sort: ", testArray)
}
