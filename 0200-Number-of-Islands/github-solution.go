// first BFS 

var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}
func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {return 0}
	cols := len(grid[0])
	count := 0	// 岛屿数
	for x:=0; x<rows; x++ {
		for y:=0; y<cols; y++ {
			if grid[x][y] == '1' {
				count++
				BFS(grid, rows, cols, x, y)
			}
		}
	}
	return count
}


// BFS，rows,cols为grid行与列数， x,y为当前起始点行、列坐标
func BFS(grid [][]byte, rows, cols, x, y int) {
	queue := make([]int, 0)
	queue = append(queue, x, y)		// 将广度优先搜索的起始点坐标加入队尾
	grid[x][y] = '2'
	for len(queue) != 0 {	// 队列非空
		curX, curY := queue[0], queue[1]  //pop栈，然后对栈进行操作,
		queue = queue[2:]		// 更新栈
		for k:=0; k<4; k++ {
			newX, newY := curX + dx[k], curY + dy[k]
			if newX>=0 && newX<rows && newY>=0 &&newY<cols && grid[newX][newY]=='1' {
				grid[newX][newY] = '2'
				queue = append(queue, newX, newY)	// 将符合该岛屿一部分的格子的坐标加入队列 (入栈)
			}
		}
	}
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
