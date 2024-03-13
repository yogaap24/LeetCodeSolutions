/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    pool := []*TreeNode{root}
    depth := 0
    
    for len(pool) > 0 {
        depth++
        level := len(pool)
        
        for i := 0; i < level; i++ {
            node := pool[0]
            pool = pool[1:]
            
            if node.Left != nil {
                pool = append(pool, node.Left)
            }
            
            if node.Right != nil {
                pool = append(pool, node.Right)
            }
        }
    }
    
    return depth
}