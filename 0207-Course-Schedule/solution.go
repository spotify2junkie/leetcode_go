// topological sort - using AOV 

func canFinish(n int, pre [][]int) bool {
	in := make([]int,n) // 指定定长
	adj := make([][]int,n) // 邻接矩阵
	next := make([]int, 0, n) // 指定容量
	for _, v := range pre {
		in[v[0]]++ // 入度 是第一个作为进入的
		adj[v[1]] = append(adj[v[1]],v[0]) // 给每一行加上自己的邻居
	}

	for i := 0; i < n; i++ {
		if in[i] == 0 {
			next = append(next,i) // 加上入度为0的 ， 这些可看成所有其他课程的prerequisite                                   // 有其他课程的prerequisite
		}
	}

	for i := 0 ; i < len(next); i++ {
		cur := next[i]
		cur_vertexs := adj[cur]
		for _, cv := range cur_vertexs {
			in[cv]--
			if in[cv] == 0 {
				next = append(next,cv) 
			}
		}
	}
	return len(next) == n 
}
