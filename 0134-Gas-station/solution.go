

func canCompleteCircuit(gas []int, cost []int) int {
	var minIdx, sum int
	min := 1<<31 - 1 //nnn bit 运用
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum < min {
			min = sum
			minIdx = i
		}
	}
	if sum < 0 {
		return -1
	}
	return (minIdx + 1) % len(gas) //nnn 不怎么喜欢用取余数
}

