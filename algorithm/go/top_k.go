package algorithm

func BotK(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	i, j, mid := l-1, r+1, nums[(l+r)/2]
	for i < j {
		for i++; nums[i] < mid; i++ {
		}
		for j--; nums[j] > mid; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if j-l+1 >= k {
		return BotK(nums, l, j, k)
	}
	return BotK(nums, j+1, r, k-(j-l+1))
}

// 从 1 开始
func TopK(nums []int, l, r, k int) int {
	return BotK(nums, l, r, r-l+1-k+1)
}
