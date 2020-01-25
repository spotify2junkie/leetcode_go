func exist(board [][]byte, word string) bool {
    for i:=0;i<len(board);i++ {
        for j:=0;j<len(board[0]);j++ {
            if search(i,j,0,board,word) {
                return true 
            }
        }   
    }
    return false 
}

func search(i, j, index int,board [][]byte,word string) bool {
    if index == len(word) {
        return true 
    }
    if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
        return false
    }
    if word[index] != board[i][j] {
        return false
    }
    temp := board[i][j]  //nnn why need this 
    board[r][c] = 0 // so no more searching if this index correct 
    if search(i-1,j,index+1,board,word) || search(i+1,j,index+1,board,word) || search(i,j+1,index+1,board,word) || search(i,j-1,index+1,board,word) { return true }
    board[i][j] = temp // if not right still need to make up this location for it's original value 
    return false 
}
