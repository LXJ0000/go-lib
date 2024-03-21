package binary_indexed_tree

type SI struct {
	tr []int
	n  int // 长度
}

func NewSI(nums ...int) *SI {
	n := len(nums)
	tr := make([]int, n+1)
	add := func(x, c int) {
		for i := x; i <= n; i += lowbit(i) {
			tr[i] += c
		}
	}
	for i := 1; i <= n; i++ {
		add(i, nums[i-1])
	}
	return &SI{
		tr: tr,
		n:  n,
	}
}

func (s *SI) Add(x, c int) {
	for i := x; i <= s.n; i += lowbit(i) {
		s.tr[i] += c
	}
}

func (s *SI) Sum(x int) int {
	sum := 0
	for i := x; i > 0; i -= lowbit(i) {
		sum += s.tr[i]
	}
	return sum
}

func (s *SI) Size() int {
	return s.n
}

func lowbit(x int) int {
	return x & -x
}

func (s *SI) Interval(l, r int) int {
	if l < 1 {
		l = 1
	}
	if r > s.Size() {
		r = s.Size()
	}
	return s.Sum(r) - s.Sum(l-1)
}
