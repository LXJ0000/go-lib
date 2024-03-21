package binary_indexed_tree

import "testing"

func TestSI(t *testing.T) {
	s := NewSI(1, 2, 1, 1, 1, 1, 1, 1, 1, 1)
	for i := 1; i < s.Size(); i++ {
		print(s.Sum(i), " ")
	}
	println()
	s.Add(1, 1)
	s.Add(3, 1)
	for i := 1; i < s.Size(); i++ {
		print(s.Sum(i), " ")
	}
	println("+====")
	println(s.Interval(1, 1))
	println(s.Interval(1, 2))
	println(s.Interval(1, 3))
	println(s.Interval(1, 4))
	println(s.Interval(2, 2))
	println(s.Interval(2, 3))
}
