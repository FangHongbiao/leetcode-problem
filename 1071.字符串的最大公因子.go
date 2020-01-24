package main

/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] 字符串的最大公因子
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	if len(str2) < len(str1) {
		str1, str2 = str2, str1
	}

	for i := len(str1); i > 0; i-- {
		if len(str1)%i != 0 || len(str2)%i != 0 {
			continue
		}
		if check(str1[:i], str1, str2) {
			return str1[:i]
		}
	}

	return ""
}

func check(t, s1, s2 string) bool {

	for i := 0; i < len(s1)/len(t); i++ {
		if s1[i*len(t):(i+1)*len(t)] != t {
			return false
		}
	}

	for i := 0; i < len(s2)/len(t); i++ {
		if s2[i*len(t):(i+1)*len(t)] != t {
			return false
		}
	}

	return true
}

func main() {
	println(gcdOfStrings("ABCABC", "ABC"))
}
// @lc code=end
