//nnn math 最大值 trick
func search(reader ArrayReader, target int) int {
    cur, end := 0, math.MaxUint32
    for cur <= end {
        mid := (cur+end) >> 1
        getted := reader.get(mid)
        if getted == math.MaxUint32 || getted > target {
            end = mid -1
        } else if getted < target {
            cur = mid +1
        } else {
            return mid
        }
    }
    return -1
}
