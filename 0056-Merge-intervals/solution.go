// slice play, array  merge and struct
package main

import (
	"fmt"
	"sort"
)

func merge(its [][]int) [][]int {
	if len(its) <= 1 {
		return its
	}
	sort.Slice(its, func(i, j int) bool { //nnn sort slice
		return its[i][0] < its[j][0]
	})
	res := make([][]int, 0, len(its))
	temp := its[0]
	for i := 1; i < len(its); i++ { //nnn start from i=1
		if its[i][0] <= temp[1] {
			temp[1] = max(temp[1], its[i][1])
		} else {
			res = append(res, temp)
			temp = its[i] //nnn 如果是第一二个对比现在就变成第三个
		}
	}
	res = append(res, temp)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := [][]int{
		{1, 4},
		{4, 5},
	}
	fmt.Println(merge(a))
}
