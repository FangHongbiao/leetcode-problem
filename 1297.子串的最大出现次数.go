package main

/*
 * @lc app=leetcode.cn id=1297 lang=golang
 *
 * [1297] 子串的最大出现次数
 */

// @lc code=start
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {

	ans := 0
	cnts := make(map[string]int)
	for i := 0; i < len(s)-minSize+1; i++ {
		t := s[i : i+minSize]
		if !isValid(t, maxLetters) {
			continue
		}
		if _, ok := cnts[t]; !ok {
			cnts[t] = 1
		} else {
			cnts[t]++
		}
	}

	for v := range cnts {
		if ans < cnts[v] {
			ans = cnts[v]
		}
	}
	return ans
}

func isValid(s string, maxLetters int) bool {
	m := make(map[byte]bool)
	for i := range s {
		m[s[i]] = true
	}

	return len(m) <= maxLetters
}

//func main() {
//	//freq := maxFreq("aaaa", 1, 3, 3)
//	//freq := maxFreq("aababcaab", 2, 3, 4)
//	//freq := maxFreq("aabcabcab", 2, 2, 3)
//	freq := maxFreq("babcbceccaaacddbdaedbadcddcbdbcbaaddbcabcccbacebda", 1, 1, 1)
//	fmt.Println(freq)
//}

// @lc code=end
