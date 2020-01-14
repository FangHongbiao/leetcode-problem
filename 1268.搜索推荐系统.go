package main
import "sort"

/*
 * @lc app=leetcode.cn id=1268 lang=golang
 *
 * [1268] 搜索推荐系统
 */

// @lc code=start
func lower_bound(st int, ed int, key string, num []string) int {
	left, right := 0, ed-1

	for left <= right {
		mid := (right + left) / 2
		if num[mid] >= key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func upper_bound(st int, ed int, key string, num []string) int {
	left, right := 0, ed-1
	for left <= right {
		mid := (right + left) / 2
		if num[mid] > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func suggestedProducts(products []string, searchWord string) [][]string {
	// 先排序
	sort.Strings(products)

	result := make([][]string, len(searchWord), len(searchWord))

	for i := 1; i <= len(searchWord); i++ {
		cur := searchWord[:i]
		lb := lower_bound(0, len(products), cur, products)
		ub := upper_bound(0, len(products), cur+"~", products)
		if lb < ub {
			if ub-lb > 3 {
				ub = lb + 3
			}
			result[i-1] = products[lb:ub]
		}
	}
	return result
}

// @lc code=end
