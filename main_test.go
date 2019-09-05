package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

var BenchaArrayLen = 10000000

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
	//fmt.Println("Regular  Quick sort: ", testArray)
}

func Test_Quick_Goroutines(t *testing.T) {
	var testArray = []int{435, 543, 435, 342, 543, 34, 532, 56594, 26, 73864, 267, 2, 1, 77, 8, 368, 686, 864, 3, 4}
	var wg sync.WaitGroup
	ch := make(chan struct{}, 4)

	QuickSortWithGoroutines(&testArray, 0, len(testArray)-1, &wg, ch)
	wg.Wait()

	for i := 0; i < len(testArray)-1; i++ {
		if testArray[i] > testArray[i+1] {
			t.Error("Not sorted!")
		}
	}
	//fmt.Println("Parallel Quick sort: ", testArray)
}

func Benchmark_bubble(b *testing.B) {
	testArray := make([]int, BenchaArrayLen)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := 0; i < BenchaArrayLen; i++ {
			testArray[i] = rand.Intn(1000)
		}
		b.StartTimer()
		BubleSort(&testArray)
	}

	for i := 0; i < len(testArray)-1; i++ {
		if testArray[i] > testArray[i+1] {
			fmt.Println("Not sorted!")
		}
	}
	//fmt.Println("Bubble sort: ", testArray)
}

func Benchmark_Quick(b *testing.B) {
	testArray := make([]int, BenchaArrayLen)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := 0; i < BenchaArrayLen; i++ {
			testArray[i] = rand.Intn(1000)
		}
		b.StartTimer()
		QuickSort(&testArray, 0, len(testArray)-1)
	}

	for i := 0; i < len(testArray)-1; i++ {
		if testArray[i] > testArray[i+1] {
			fmt.Println("Not sorted!")
		}
	}
	//fmt.Println("Quick sort: ", testArray)
}

func Benchmark_Quick_Goroutines(b *testing.B) {

	testArray := make([]int, BenchaArrayLen)
	var wg sync.WaitGroup
	ch := make(chan struct{}, runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := 0; i < BenchaArrayLen; i++ {
			testArray[i] = rand.Intn(1000)
		}
		b.StartTimer()
		QuickSortWithGoroutines(&testArray, 0, len(testArray)-1, &wg, ch)
		wg.Wait()
	}

	for i := 0; i < len(testArray)-1; i++ {
		if testArray[i] > testArray[i+1] {
			fmt.Println("Not sorted!")
		}
	}
	//fmt.Println("Quick sort: ", testArray)
}
