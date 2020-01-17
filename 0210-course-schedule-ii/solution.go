//nnn topological sort 
func findOrder(n int, prerequisites [][]int) []int {
    in := make([]int,n) // 指定定长
	adj := make([][]int,n) // 邻接矩阵
	next := make([]int, 0, n) // 指定容量
    for _,v := range prerequisites {
        in[v[0]]++
        adj[v[1]] = append(adj[v[1]],v[0]) // 给每一行加上自己的邻居
    }

    for i:=0;i<n;i++ {
        if in[i] == 0 {
            next = append(next,i)
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
    if len(next)!= n { //nnn problem 207 翻版
        return []int{}
    }
	return next  
}