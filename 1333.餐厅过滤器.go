import "sort"

/*
 * @lc app=leetcode.cn id=1333 lang=golang
 *
 * [1333] 餐厅过滤器
 */

// @lc code=start
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {

	// 对餐馆按照评分排序
	// 排序规则: 按照 rating 从高到低排序。如果 rating 相同，那么按 id 从高到低排序
	sort.Slice(restaurants, func(i, j int) bool {
		if restaurants[j][1] != restaurants[i][1] {
			return restaurants[i][1] > restaurants[j][1]
		}
		return restaurants[i][0] > restaurants[j][0]
	})

	ans := make([]int, 0)
	for _, r := range restaurants {
		if veganFriendly == 1 && r[2] != 1 {
			continue
		}
		if r[3] <= maxPrice && r[4] <= maxDistance {
			ans = append(ans, r[0])
		}
	}
	return ans
}

// @lc code=end
