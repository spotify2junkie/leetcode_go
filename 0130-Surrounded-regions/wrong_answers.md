# 错误解答
例子
```
[["X","X","X","X"]
["X","O","O","X"],
["X","X","O","X"],
["X","O","X","X"]]
```
错误输出
```
[["X","X","X","X"],
["X","O","O","X"],
["X","X","O","X"],
["X","O","X","X"]]
```

正确输出
```
[["X","X","X","X"],
["X","X","X","X"],
["X","X","X","X"],
["X","O","X","X"]]
```
# 错误代码 ！！
```
var dx = [4]int{-1, 0, 1, 0} // moving directions
var dy = [4]int{0, 1, 0, -1} // moving directions 

func solve(board [][]byte) {
    for i := range board {
        for j := range board[i] {
            if i != 0 && j != 0 && i != len(board)-1 && j != len(board[i])-1 {
                if board[i][j] == 'O' { //nnn string manipulation 
                    dfs(i,j,board)
                }
            }
        }
    }
    return 
}

func dfs(i, j int, board [][]byte) {  ////nnn 相比200
	count := 0 
	for k:=0; k<4; k++ {
		newX,newY := i+dx[k],j+dy[k]
		if board[newX][newY] == 'X' {
			count += 1
		}
	}
	if count == 4 {
		board[i][j] = 'X'
	}
    return 
} 
```

# 错误思路
- 我想的是必须被'X' 包围，但是其实只要边界不为’O'就行
