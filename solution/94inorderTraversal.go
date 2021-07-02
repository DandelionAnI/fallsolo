package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	inorder(root, &ans)
	return ans
}

func inorder(root *TreeNode, ans *[]int) {
	if root != nil {
		inorder(root.Left, ans)
		*ans = append(*ans, root.Val)
		inorder(root.Right, ans)
	}
}
