package main

import (
	"fmt"
	"math"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if left >= root.Val || right <= root.Val {
		return false
	}
	return isBST(root.Left, left, root.Val) && isBST(root.Right, root.Val, right)
}

func main() {
	tcs := []struct {
		pre, in []int
		ans     bool
	}{

		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			false,
		},

		{
			[]int{},
			[]int{},
			true,
		},

		{
			[]int{2, 1, 3},
			[]int{1, 2, 3},
			true,
		},

		{
			[]int{10, 5, 15, 6, 20},
			[]int{5, 10, 6, 15, 20},
			false,
		},

		{
			[]int{3, 2, 1},
			[]int{2, 3, 1},
			false,
		},
		// 可以多个 testcase
	}
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		if tc.ans == isValidBST(kit.PreIn2Tree(tc.pre, tc.in)) {
			fmt.Println("Test Passed")
		}
	}
}
