//nnn slice , array 
func longestConsecutive(nums []int) int {
	count := make(map[int]int)
	res := 0
	for _, i := range nums {
		_, ok := count[i]
        if !ok { //nnn didnt see before 
            left := count[i-1]
			right := count[i+1]
			sum := left + right + 1
            if sum > res {
                res = sum 
            }
            count[i],count[i+right],count[i-left] = sum, sum, sum  //nnn 绝了，直接操作比自己大的count
            // 假如right == 1 and i = 3,在4的时候count[3] = 1 ,然后到3的时候right = 1 ,res 和sum 都为 2,然后 
            // count[i+right] = 2 cuz count[i+1] = sum 
        }
    }
    return res 
}
