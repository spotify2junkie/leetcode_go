//nnn slice array

func canCompleteCircuit(gas []int, cost []int) int {
	var minIdx, sum int
	min := 1<<31 - 1 //nnn bit 运用
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum < min { // start_idx
			min = sum
			minIdx = i // 记录min_idx
		}
	}
	if sum < 0 {
		return -1
	}
	return (minIdx + 1) % len(gas) //nnn 不怎么喜欢用取余数
}

