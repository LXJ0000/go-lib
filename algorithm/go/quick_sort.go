package algorithm

func QuickSort(nums []int, l, r int) {
	if l == r {
		return
	}
	i, j, mid := l-1, r+1, nums[(l+r)>>1]
	for i < j {
		for i++; nums[i] < mid; i++ {
		}
		for j--; nums[j] > mid; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	QuickSort(nums, l, j)
	QuickSort(nums, j+1, r)
}

