import "sort"

/*
 * @lc app=leetcode.cn id=1344 lang=golang
 *
 * [1344] 跳跃游戏 V
 */

// @lc code=start

// func maxJumps(arr []int, d int) int {
// 	max := 1
// 	for i := range arr {
// 		cur := walk(arr, i, d)
// 		if cur > max {
// 			max = cur
// 		}
// 	}
// 	return max
// }
// func walk(arr []int, idx int, d int) int {
// 	max := 1
// 	// 向右跳跃
// 	for i := idx + 1; i < minInts(len(arr), idx+d+1) && arr[idx] > arr[i]; i++ {
// 		cur := walk(arr, i, d) + 1
// 		if cur > max {
// 			max = cur
// 		}
// 	}
// 	// 向左跳跃
// 	for i := idx - 1; i >= maxInts(0, idx-d) && arr[idx] > arr[i]; i-- {
// 		cur := walk(arr, i, d) + 1
// 		if cur > max {
// 			max = cur
// 		}
// 	}
// 	return max
// }

// var cache []int

// func maxJumps(arr []int, d int) int {
// 	// 初始化缓存
// 	cache = make([]int, len(arr))
// 	max := 1
// 	for i := range arr {
// 		cur := walk(arr, i, d)
// 		if cur > max {
// 			max = cur
// 		}
// 	}
// 	return max
// }
// func walk(arr []int, idx int, d int) int {

// 	if cache[idx] != 0 {
// 		return cache[idx]
// 	}
// 	max := 1
// 	// 向右跳跃
// 	for i := idx + 1; i < minInts(len(arr), idx+d+1) && arr[idx] > arr[i]; i++ {
// 		cur := walk(arr, i, d) + 1
// 		if cur > max {
// 			max = cur
// 		}
// 	}
// 	// 向左跳跃
// 	for i := idx - 1; i >= maxInts(0, idx-d) && arr[idx] > arr[i]; i-- {
// 		cur := walk(arr, i, d) + 1
// 		if cur > max {
// 			max = cur
// 		}
// 	}
// 	cache[idx] = max
// 	return max
// }

func maxJumps(arr []int, d int) int {

	// 先进行排序, 注意, 我们需要的是排序的索引, 索引先把数值,索引组合起来
	indexs := make([][]int, len(arr))
	for i, e := range arr {
		indexs[i] = make([]int, 2)
		indexs[i][0] = e
		indexs[i][1] = i
	}

	// 排序
	sort.Slice(indexs, func(i, j int) bool {
		return indexs[i][0] < indexs[j][0]
	})

	max := 1
	dp := make([]int, len(arr))

	for _, e := range indexs {
		idx := e[1]
		dp[idx] = 1
		// 向右跳跃
		for i := idx + 1; i < minInts(len(arr), idx+d+1) && arr[idx] > arr[i]; i++ {
			dp[idx] = maxInts(dp[idx], dp[i]+1)
		}
		// 向左跳跃
		for i := idx - 1; i >= maxInts(0, idx-d) && arr[idx] > arr[i]; i-- {
			dp[idx] = maxInts(dp[idx], dp[i]+1)
		}
		max = maxInts(dp[idx], max)
	}

	return max
}

func maxInts(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minInts(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end
