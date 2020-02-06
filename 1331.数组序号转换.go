import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1331 lang=golang
 *
 * [1331] 数组序号转换
 */

// @lc code=start
func arrayRankTransform(arr []int) []int {

	// 创建一个map， 用于去重，同时用于记录元素对应的位置
	ele2Idx := make(map[int]int, len(arr))

	// 用于去重，位置默认
	for _, e := range arr {
		ele2Idx[e] = -1
	}

	// 对去重元素进行排序
	sortedArr := make([]int, 0, len(arr))
	for key := range ele2Idx {
		sortedArr = append(sortedArr, key)
	}
	sort.Ints(sortedArr)

	// 根据排序数组记录元素应该在的位置，存放到map中
	for i, e := range sortedArr {
		ele2Idx[e] = i + 1
	}

	// 将原始数据原地替换为位置，返回
	for i := 0; i < len(arr); i++ {
		arr[i] = ele2Idx[arr[i]]
	}

	return arr
}

// @lc code=end
