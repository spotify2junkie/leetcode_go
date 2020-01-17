// hash table 

func dfs (i,j,lenR,lenC int,board[][]byte) {
    if i< 0 || j<0 || i>=lenR || j>=lenC || board[i][j] == 'X' || board[i][j] == 'G' {return }
    board[i][j] = 'G'
    dfs(i+1,j,lenR,lenC,board)
    dfs(i-1,j,lenR,lenC,board)
    dfs(i,j+1,lenR,lenC,board)
    dfs(i,j-1,lenR,lenC,board)
}


func solve(board [][]byte) {
	lenR := len(board)
    if board == nil || lenR == 0 {
        return 
    }
    lenC := len(board[0])
    if lenC == 0 {
        return 
    }
	// 先搜索并标记
	for i:=0;i<lenR;i++ {
        dfs(i,0,lenR,lenC,board)
        dfs(i,lenC-1,lenR,lenC,board)
    }
    for i:=0;i<lenC;i++ {
        dfs(0,i,lenR,lenC,board)
        dfs(lenR-1,i,lenR,lenC,board)
    }
    replaceMap := map[byte]byte{'G':'O','O':'X','X':'X'} //nnn map[type]type  
    for i:=0;i<lenR;i++ {
        for j:=0;j<lenC;j++ {
            board[i][j] = replaceMap[board[i][j]] //nnn amazing hash table 
        }
    }
}
