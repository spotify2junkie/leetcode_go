// 递归的方法
// 
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return ish(root.Left,root.Right) // return bool 
}

func ish(l, r *TreeNode) bool {   
    if l == nil || r == nil {
        return false
    }
    
    if l.Val == r.Val {
        return true
    }
    
    return ish(l.Left,l.Right) && ish(r.Left,r.Right)
}