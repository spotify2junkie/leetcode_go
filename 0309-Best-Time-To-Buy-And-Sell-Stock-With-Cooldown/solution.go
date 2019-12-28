// dynamic programming


func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	buy := make([]int,n+1) // 这样才能取slice 
	sel := make([]int,n+1)
	buy[1] = 0 - prices[0]

	for i:=2; i <= n; i++ {
		buy[i] = max(buy[i-1],sel[i-2]-prices[i-1]) // 代表的都是profit
		sel[i] = max(sel[i-1],buy[i-1]+prices[i-1]) // 代表的都是profit 
	}
	return sel[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b 
}
