/*
 * @lc app=leetcode.cn id=1325 lang=golang
 *
 * [1325] 删除给定值的叶子节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {

	if root == nil {
		return nil
	}

	// 把root的左孩子赋值为处理后的
	root.Left = removeLeafNodes(root.Left, target)
	// 把root的右孩子赋值为处理后的
	root.Right = removeLeafNodes(root.Right, target)

	// 如果当前节点满足左右孩子为空并且值为target，则删除，返回nil
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	// 否则返回root， 注意此时它的左右孩子已经处理好
	return root
}

// @lc code=end
