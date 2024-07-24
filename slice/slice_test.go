package slice

import (
	"fmt"
	"testing"

	"github.com/LXJ0000/go-lib"
	"github.com/LXJ0000/go-lib/errs"
	"github.com/stretchr/testify/require"
)

// go get "github.com/stretchr/testify/require"

func TestMax(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: -1,
		},
		{
			name:     "mixed numbers",
			nums:     []int{-10, 0, 10, -5, 5},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Max(tc.nums...)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: -5,
		},
		{
			name:     "mixed numbers",
			nums:     []int{-10, 0, 10, -5, 5},
			expected: -10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Min(tc.nums...)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: -15,
		},
		{
			name:     "mixed numbers",
			nums:     []int{-10, 0, 10, -5, 5},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.nums...)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestDeleteAt(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		index    int
		expected []int
		err      error
	}{
		{
			name:     "positive",
			nums:     []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
			err:      nil,
		},
		{
			name:     "negative",
			nums:     []int{-5, -4, -3, -2, -1},
			index:    -1,
			expected: nil,
			err:      errs.NewErrIndexOutOfRange(5, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := DeleteAt(tc.nums, tc.index)
			require.Equal(t, tc.expected, result)
			require.Equal(t, tc.err, err)
		})
	}
}

func TestUnique(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "positive",
			nums:     []int{1, 2, 2, 2, 3, 5, 4, 5, 6, 6, 1, 3, 4, 5},
			expected: []int{1, 2, 3, 5, 4, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Unique(tc.nums)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestMap(t *testing.T) {
	src := []int{1, 2, 3, 4}
	dst := Map[int, int](src, func(src int) int {
		return src + 1
	})
	fmt.Println(dst)
}

// TODO DeleteAt benchmark
func BenchmarkDeleteAt(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}
	index := 500

	for i := 0; i < b.N; i++ {
		_, err := DeleteAt(nums, index)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDeleteAtV2(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}
	index := 500

	for i := 0; i < b.N; i++ {
		_, err := DeleteAtV2(nums, index)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func DeleteAtV2[T lib.Ordered](nums []T, index int) ([]T, error) {
	if index < 0 || index >= len(nums) {
		return nil, errs.NewErrIndexOutOfRange(len(nums), index)
	}
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	nums = nums[:len(nums)-1]
	return Shrink(nums), nil
	//	BenchmarkDeleteAt-8             26257840                45.50 ns/op
	//  BenchmarkDeleteAtV2-8            4379086               263.8 ns/op
}
