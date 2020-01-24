package main

/*
 * @lc app=leetcode.cn id=915 lang=golang
 *
 * [915] 分割数组
 */

// @lc code=start
func partitionDisjoint(A []int) int {
	l := len(A)
	maxLs := make([]int, l)
	minRs := make([]int, l)

	maxLs[0] = A[0]
	minRs[l-1] = A[l-1]
	for i := 1; i < l; i++ {
		if A[i] > maxLs[i-1] {
			maxLs[i] = A[i]
		} else {
			maxLs[i] = maxLs[i-1]
		}
	}

	for i := l - 2; i >= 0; i-- {
		if A[i] < minRs[i+1] {
			minRs[i] = A[i]
		} else {
			minRs[i] = minRs[i+1]
		}
	}

	for i := 0; i < l-1; i++ {
		if maxLs[i] <= minRs[i+1] {
			return i + 1
		}
	}
	return 0
}

// @lc code=end
