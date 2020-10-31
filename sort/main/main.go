package main

import (
	"base/sort/main/sort"
	"fmt"
	"math/rand"
)

func main() {
	s := make([]int, 0, 16)
	for i := 0; i < 16; i++ {
		s = append(s, rand.Intn(100))
	}
	fmt.Println(s)
	sort.QuickSort(s)
	// s = sort.MergeSort(s)
	// sort.HeapSort(s)
	fmt.Println(s)
}
