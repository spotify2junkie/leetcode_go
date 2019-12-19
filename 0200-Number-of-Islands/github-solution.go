package problem0200

func numIslands(grid [][]byte) int {
	// 获取 grid 的大小
	m := len(grid) // number of column
	if m == 0 {
		return 0 //return if zero size
	}
	n := len(grid[0])

	// bfs 搜索时，存放点坐标的队列
	x := make([]int, 0, m*n) // make 存储, m*n 为容量值
	y := make([]int, 0, m*n)

	// 往队列中添加 (i,j) 点，并修改 (i,j) 点的值
	// 避免重复搜索
	var add = func(i, j int) {
		x = append(x, i) // 局部
		y = append(y, j)
		grid[i][j] = '0'
	}

	// 从坐标队列中，取出坐标点
	var pop = func() (int, int) {
		i := x[0]
		x = x[1:]
		j := y[0]
		y = y[1:]
		return i, j
	}

	var bfs = func(i, j int) int {
		if grid[i][j] == '0' {
			return 0
		}

		add(i, j)

		for len(x) > 0 {
			i, j = pop()

			// 搜索 (i,j) 点的 上下左右 四个方位
			if 0 <= i-1 && grid[i-1][j] == '1' {
				add(i-1, j)
			}

			if 0 <= j-1 && grid[i][j-1] == '1' {
				add(i, j-1)
			}

			if i+1 < m && grid[i+1][j] == '1' {
				add(i+1, j)
			}

			if j+1 < n && grid[i][j+1] == '1' {
				add(i, j+1)
			}
		}

		return 1
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += bfs(i, j)
		}
	}

	return res
}

// another one ---
// DFS

var dx = [4]int{-1, 0, 1, 0} // 控制遍历方向
var dy = [4]int{0, -1, 0, 1}

func judge(grid [][]byte, x, y int) bool {
	// 不能走但反过来
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) { //边界学习了
		return true
	}
	if grid[x][y] == '0' {
		return true
	}
	return false

}

func dfs(grid [][]byte, x, y int) { // 递归
	if judge(grid, x, y) { // 如果不能走return
		return
	}
	grid[x][y] = '0'
	for i := 0; i < 4; i++ { // 遍历当前位置的四个方向，并累计与这四个方向连接的土地面积
		dfs(grid, x+dx[i], y+dy[i]) // 这个真的很妙
	}
}

func numIslands(grid [][]byte) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				cnt += 1
			}
		}
	}
	return cnt
}
