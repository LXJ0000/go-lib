package algorithm

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{5, 4, 3, 1, 2, 3, 4, 9, 6, 3, 6, 3, 2, 7, 1000}
	QuickSort(nums, 0, len(nums)-1)
	log.Fatalln(nums)
}
