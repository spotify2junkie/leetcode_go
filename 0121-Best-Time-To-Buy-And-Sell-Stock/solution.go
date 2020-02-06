// dp  dynamic programming 
func maxProfit(prices []int) int {
    max := 0
    temp := 0
    for i := 1; i < len(prices);i ++ {
        temp += prices[i]-prices[i-1]
        if temp > max {
            max = temp
        } else if temp < 0 {
            temp = 0 
        }
    }
    return max 
}
