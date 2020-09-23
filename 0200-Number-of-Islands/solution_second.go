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
	for i := 0; i < 4; i++ { // 遍历当前位置的四个方向，
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

