
// copied solution for learning
// original ansnwer https://github.com/halfrost/LeetCode-Go/blob/master/template/UnionFind.go
// Init define

type UnionFind struct {
	parent, rank []int
	count        int
}

func (uf *UnionFind) Init(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

// Find define
func (uf *UnionFind) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	// compress path
	for p != uf.parent[p] {
		tmp := uf.parent[p]
		uf.parent[p] = root
		p = tmp
	}
	return root
}

// Union define
func (uf *UnionFind) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return // 解决了自己的parent 是自己，以及其他人共享parent 这个问题
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parent[proot] = qroot
	} else {
		uf.parent[qroot] = proot
		if uf.rank[proot] == uf.rank[qroot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}

func (uf *UnionFind) TotalCount() int {
	return uf.count
}

// main solution
func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}
	uf := UnionFind{}
	uf.Init(n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.TotalCount()
}