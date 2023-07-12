/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSameTree(p *TreeNode, q *TreeNode) bool {

    // If either p or q is not empty, check if p is equal to q
    if p == nil || q == nil {
        return p == q
    
    // If the current node for p is equal to the current node for q, check if the
    // left and right subtrees are equal
    } else if p.Val == q.Val {
        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) 
    
    // Else p and q are not equal, return false
    } else {
        return false
    }   
}