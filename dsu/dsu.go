package dsu

type Dsu struct {
	p    []int
	size []int
}

func NewDsu(n int) *Dsu {
	p := make([]int, 0, n)
	size := make([]int, 0, n)
	for i := 0; i < n; i++ {
		p[i] = i
		size[i] = 1
	}
	return &Dsu{p: p, size: size}
}

// Find 路径压缩
func (d *Dsu) Find(x int) int {
	if x != d.p[x] {
		d.p[x] = d.Find(d.p[x])
	}
	return d.p[x]
}

// Merge 启发式合并
func (d *Dsu) Merge(x, y int) {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return
	}
	if d.size[px] < d.size[py] {
		px, py = py, px // swap
	}
	d.p[py] = px
	d.size[px] += d.size[py]
}
