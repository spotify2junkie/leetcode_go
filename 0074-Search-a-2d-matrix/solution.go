// binary search 
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false 
    }
    n := len(matrix[0])
    if n == 0 {
        return false 
    }

    if target < mat[0][0] || mat[m-1][n-1] < target {
        return false 
    }

    // find the `right` row 
    r := 0 
    for r<m && m[r][0] < target {
        r++
    }
    r--  // get the right row that match r<target 
    // binary search 
    i,j := 0,n-1
    for i<=j {
        med := (i+j)>>1
        switch {
        case mat[r][med] < target:
            i=med+1
        case mat[r][med]>target:
            j=med-1
        default:
            return true 
        }
    }
    return mat[r][j] == target 
}