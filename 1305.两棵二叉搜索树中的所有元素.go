import "fmt"

/*
 * @lc app=leetcode.cn id=1305 lang=golang
 *
 * [1305] 两棵二叉搜索树中的所有元素
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
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := inOrderTravel(root1)
	arr2 := inOrderTravel(root2)
	fmt.Println(arr1, arr2)
	ans := merge(arr1, arr2)
	return ans
}

// inOrderTravel 返回中序遍历序列
func inOrderTravel(root *TreeNode) []int {

	if root == nil {
		return make([]int, 0)
	}

	left := inOrderTravel(root.Left)
	right := inOrderTravel(root.Right)

	return append(append(left, root.Val), right...)
}

// merge 合并两个有序数组
func merge(left, right []int) []int {
	ans := make([]int, len(left)+len(right))

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ans[k] = left[i]
			i++
		} else {
			ans[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		ans[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		ans[k] = right[j]
		k++
		j++
	}
	return ans
}

// @lc code=end
