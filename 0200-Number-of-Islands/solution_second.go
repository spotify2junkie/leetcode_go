// another one ---
// DFS
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
}