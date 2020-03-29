// BFS from 2-D
func maxDistance(grid [][]int) int {
    // zero: ocean; one: land 
    // step 1: mark all land 
    // step 2: mark all ocean 
    // step 3: 笛卡尔积
    // -- dumb solution 
    dx := []int{0, 0, 1, -1}
    dy := []int{1, -1, 0, 0}
    m,n := len(grid), len(grid[0])
    if m == 0 || n == 0 {
        return -1
    }
    queue := [][]int{}
    for i:=0;i<m; i++ {
        for j:=0;j< n; j++ {
            if grid[i][j] == 1 {
                queue = append(queue, []int{i,j})
            }
        }
    }
    if len(queue) == 0 || len(queue) == len(grid)*len(grid[0]) {
        return -1
    }
    ans := 0 

    for len(queue) != 0 {
        ans++ //nnn this wors for manhattan distance 
        tmp := queue
        queue = nil 
        for len(tmp) > 0 {
            cur := tmp[0] //nnn push from 头
            tmp = tmp[1:]
            curX, curY := cur[0], cur[1]
            for i:= 0; i < 4;i++ {
                newX := curX+dx[i]
                newY := curY+dy[i]
                if newX < 0 || newX >= m || newY < 0 || newY >= n || grid[newX][newY] != 0 {
                    continue
                }
                queue = append(queue, []int{newX, newY})
                grid[newX][newY] = 2 // 代表以及被遍历过了
            }
        }
    }
    return ans-1
}
  



