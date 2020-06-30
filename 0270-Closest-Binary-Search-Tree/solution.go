//nnn binary search and tree 
func closestValue(root *TreeNode, target float64) int {
    closest := root.Val //nnn initiate  
    for root!= nil{
        if math.Abs(float64(root.Val) - target) < math.Abs(float64(closest) - target){
            closest = root.Val
        }
        
        if target > float64(root.Val){
            root = root.Right
        } else{
            root = root.Left
        }
    }
    
    return closest
}
