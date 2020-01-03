// two pointers 

// func maxArea(height []int) int {
//     area := 0
//     maxOfArea := 0 
//     if len(height) == 2  && height[1] >= height[0] {
//         return 1*height[0]
//     }
//     for i:=1;i < len(height);i++ {
//         area = helper(height, i)
//         maxOfArea = max(maxOfArea,area)
//     }
//     return maxOfArea
// }


// func helper(height []int, index int) int {
//     maxA := 0
//     for i:=0;i < index;i++ {
//         if height[i] >= height[index] {
//             area := height[index] * (index-i)
//             maxA = max(maxA,area)
//         }
//     }

//     for i:=len(height)-1;i > index;i-- {
//         if height[i] >= height[index] {
//             area := height[index] * (i-index)
//             maxA = max(maxA,area)
//         }
//     }     
//     return maxA
// }



// func max(i, j int) int {
//     if i > j {
//         return i 
//     } else {
//         return j 
//     }
// }

func maxArea(height []int) int {
    i,j := 0, len(height)-1 // 两个指针
    max := 0 
    for i != j {
        left, right := height[i],height[j]
        h := min(height[i],height[j])
        area := h * (j-i) // 计算现在的面积
        max = maxc(area,max)

        if left < right {
            i++  // 主要在于指针移动的条件
        } else { j--}
    }
    return max
}



func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a 
    }
}

func maxc(a, b int) int {
    if a > b {
        return a
    } else {
        return b 
    }
}