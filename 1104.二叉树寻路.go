package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=1104 lang=golang
 *
 * [1104] 二叉树寻路
 */

// @lc code=start
func pathInZigZagTree(label int) []int {

	cnt := 1
	row := 0
	for cnt <= label {
		cnt *= 2
		row++
	}

	if row%2 == 1 {
		return findPath(row, label)
	} else {
		return findPath(row, getSymmetry(label, row))
	}

}

// 搜索路径
func findPath(row int, label int) []int {
	if row == 1 {
		return []int{1}
	}

	var cur int
	if row%2 == 1 {
		cur = label
	} else {
		cur = getSymmetry(label, row)

	}
	remain := findPath(row-1, label/2)
	return append(remain, cur)
}

// 获得对称位置，因为完全二叉树可以计算每一行开始与结束位置，可以计算相对距离
func getSymmetry(label, row int) int {
	start, end := (int)(math.Pow(2, (float64)(row)-1)), int(math.Pow(2, (float64)(row))-1)
	diff := label - start
	return end - diff
}

// @lc code=end
