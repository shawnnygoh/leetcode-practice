/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Main Idea: Find the height of the left and the right subtrees of a node recursively and assign height to that node as the maximum of the heights of its two children plus 1.
// Time Complexity: O(n)
// Space Complexity: O(n) due to recursive stack. 

// Create `max` function which returns the larger integer between integers `x` and `y`. 
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func maxDepth(root *TreeNode) int {

    // If the tree is empty, return 0.
    if root == nil {
        return 0

    // Else, return the maximum value of the depths of the left and right roots and add 1 to it for the current node.
    } else {
        return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
    }
}