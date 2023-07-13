/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    // Create a variable called `inorder` that is assigned to and an empty integer
    // array 
    var inorder = []int{}

    // If the root of the tree is empty or null, return the value of `inorder` 
    if root == nil {
        return inorder
    // Else, append the root value to the inorder traversal of the left tree node
    // and store it in `inorder`
    } else {
        inorder = append(inorderTraversal(root.Left), root.Val)

        // Then, append the inorder traversal of the right tree node to `inorder` and
        // return it
        return append(inorder, inorderTraversal(root.Right)...)
    }
}