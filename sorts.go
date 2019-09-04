package main

//BubleSort classic
func BubleSort(input *[]int) {
	var buff int
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(*input)-1; i++ {
			if (*input)[i] > (*input)[i+1] {
				buff = (*input)[i]
				(*input)[i] = (*input)[i+1]
				(*input)[i+1] = buff
				sorted = false
			}
		}
	}
}

//QuickSort simple
func QuickSort(input *[]int, left int, right int) {
	var divideValue = (*input)[(left+right)/2]
	var i = left
	var j = right
	var buff int

	for i <= j {
		for (*input)[i] < divideValue {
			i++
		}
		for (*input)[j] > divideValue {
			j--
		}
		if i <= j {
			buff = (*input)[i]
			(*input)[i] = (*input)[j]
			(*input)[j] = buff
			i++
			j--
		}
	}
	if left < j {
		QuickSort(input, left, j)
	}
	if i < right {
		QuickSort(input, i, right)
	}
}

func main() {

}
