func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    if len(matrix[0]) == 0 {
        return []int{}
    }
    res := []int{}
    top,right,bottom,left := 0,len(matrix[0])-1,len(matrix)-1,0
    for top <= bottom && left <= right {
        for i:= left;i<=right;i++ {
            res = append(res,matrix[top][i])
        }
		top++ 
		//nnn traverse down 
        for i := top;i<= bottom;i++ {
            res = append(res,matrix[i][right])
        }
		right--
		//nnn traver left 
        for i := right; i >= left; i-- {
			if top > bottom { //nnn boundary
				continue
			}
			res = append(res, matrix[bottom][i])
		}
		bottom--
		// traverse up
		for i := bottom; i >= top; i-- {
			if left > right { //nnn boundary 
				continue
			}
			res = append(res, matrix[i][left])
		}
		left++
    }
    return res 
}