import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1329 lang=golang
 *
 * [1329] 将矩阵按对角线排序
 */

// @lc code=start
func diagonalSort(mat [][]int) [][]int {

	m := len(mat)
	n := len(mat[0])

	x1, y1 := m-1, 0
	x2, y2 := m-1, 0

	for x1 >= 0 && y1 < n {
		arr := getOne(mat, m, n, x1, y1, x2, y2)
		sort.Ints(arr)
		reAssign(mat, m, n, x1, y1, x2, y2, arr)
		x1--
		y2++
		if y2 == n {
			x2--
			y2--
		}

		if x1 < 0 {
			x1++
			y1++
		}

	}

	return mat
}

// getOne给定左上角左边(x1, y1) 和 右下角左边(x2, y2)， 得到这条对角线元素
func getOne(mat [][]int, m int, n int, x1 int, y1 int, x2 int, y2 int) []int {

	fmt.Println(x1, y1, x2, y2)
	arr := make([]int, 0, x2-x1+1)

	for x2 >= x1 {
		arr = append(arr, mat[x1][y1])
		x1++
		y1++
	}

	return arr
}

// reAssign 为对角线重新赋值(按照从左上到右下，元素在arr中)
func reAssign(mat [][]int, m int, n int, x1 int, y1 int, x2 int, y2 int, arr []int) {

	for _, e := range arr {
		mat[x1][y1] = e
		x1++
		y1++
	}

}

// @lc code=end
