/*
 * @lc app=leetcode.cn id=1317 lang=golang
 *
 * [1317] 将整数转换为两个无零整数的和
 */

// @lc code=start
func getNoZeroIntegers(n int) []int {

	// 枚举所有可能
	for i := 1; i <= n/2; i++ {
		if noZero(i) && noZero(n-i) {
			return []int{i, n - i}
		}
	}
	return []int{}
}

// noZero 判断 num 中是否含有0， 不含返回true， 否则返回false
func noZero(num int) bool {

	for num != 0 {
		if num%10 == 0 {
			return false
		}
		num /= 10
	}

	return true
}

// @lc code=end
