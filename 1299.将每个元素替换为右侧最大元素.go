/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {

	// 构建辅助数组
	help := make([]int, len(arr))
	help[len(arr)-1] = arr[len(arr)-1]

	for i := len(help) - 2; i >= 0; i-- {
		if arr[i] < help[i+1] {
			help[i] = help[i+1]
		} else {
			help[i] = arr[i]
		}
	}

	// 向右挪一位就是解
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = help[i+1]
	}
	arr[len(arr)-1] = -1
	return arr
}

// @lc code=end
