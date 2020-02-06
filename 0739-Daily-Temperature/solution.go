func dailyTemperatures(T []int) []int {
    days := []int{}
    for index, temp := range T {
        if temp < 100 && index != len(T) - 1 {
            cur := max(temp, T[index+1:])
            if cur.max > temp {
                days = append(days,cur.loc + 1)
            } else {
                days = append(days, 0) // 如果没有一个大的
            }
		} else { days = append(days, 0)} // 尾端节点打分
		
    }
    return days 
}
        



type item struct {
    loc int
    max int
}


func max (cur int, stack []int) item {
    a := item{0,cur}
    for index, val := range stack { 
        if val > a.max {
            a.max = val
            a.loc = index
            break
        }  
    }
    return a
}
