package main

import (
	"fmt"

	binarysearch "github.com/barnabasSol/DS/binary_search"
)

func main() {
	randomNumbers := []int{3, 7, 12, 25, 34, 45, 56, 67, 78, 89, 100}
	result := binarysearch.BinarySearch(randomNumbers, 25)
	fmt.Println(result)

}
