//nnn Binary Search Tree
// 这道题就是找到中间的点作为root node , root node左边的是left tree ....
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return &TreeNode{Val: nums[len(nums)/2], Left: sortedArrayToBST(nums[:len(nums)/2]),
                    Right: sortedArrayToBST(nums[len(nums)/2 + 1:])}
}
