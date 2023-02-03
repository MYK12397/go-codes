package main

import (
	"fmt"
)

//bubble sort
func bubblesort(ar []int) {
	for i := 0; i < len(ar); i++ {
		swapped := false
		for j := i + 1; j < len(ar)-1; j++ {

			if ar[j] < ar[i] {
				ar[i], ar[j] = ar[j], ar[i]
				swapped = true
			}

			if !swapped {
				break
			}
		}
	}
	fmt.Println(ar)
}

//insertion sort
func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {

		key := ar[i]
		j := i - 1

		for j >= 0 && ar[j] > key {
			ar[j+1] = ar[j]
			j = j - 1
		}
		ar[j+1] = key
	}
	fmt.Println(ar)
}

//quick sort
func quicksort(ar []int) []int {

	if len(ar) < 2 {
		return ar
	}

	left, right := 0, len(ar)-1

	pivot := left
	ar[pivot], ar[right] = ar[right], ar[pivot]

	for i, _ := range ar {

		if ar[i] < ar[right] {
			ar[left], ar[i] = ar[i], ar[left]
			left++
		}
	}
	ar[left], ar[right] = ar[right], ar[left]
	quicksort(ar[:left])
	quicksort(ar[left+1:])

	return ar
}

//selection sort
func selectionsort(a []int) {
	n := len(a)

	for i := 0; i < n; i++ {
		minIdx := i

		for j := i; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
}

func main() {
	ar := []int{3, -2, 40, 1, 9, 5}

	insertionSort(ar)
	fmt.Println(ar)
	quicksort(ar)
	selectionsort(ar)
	fmt.Println(ar)

}
