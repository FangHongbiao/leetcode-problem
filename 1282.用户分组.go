package main

/*
 * @lc app=leetcode.cn id=1282 lang=golang
 *
 * [1282] 用户分组
 */

// @lc code=start
func groupThePeople(groupSizes []int) [][]int {

	cnts := make(map[int][]int)

	for i, groupSize := range groupSizes {
		if _, ok := cnts[groupSize]; !ok {
			cnts[groupSize] = make([]int, 0)
		}
		cnts[groupSize] = append(cnts[groupSize], i)
	}

	result := make([][]int, 0)

	for k, v := range cnts {
		sz := len(v)
		for j := 0; j < sz/k; j++ {
			result = append(result, v[j*k:(j+1)*(k)])
		}
	}
	return result
}

// @lc code=end
