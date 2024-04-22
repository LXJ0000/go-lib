package algorithm

import (
	"log"
	"testing"
)

func TestTopK(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	log.Fatalln(TopK(nums, 0, len(nums)-1, 7))
}
