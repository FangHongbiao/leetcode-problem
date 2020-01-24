package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=1309 lang=golang
 *
 * [1309] 解码字母到整数映射
 */

// @lc code=start
func freqAlphabets(s string) string {
	return process(0, s)
}

func process(idx int, s string) string {

	if idx >= len(s) {
		return ""
	}

	if idx+2 < len(s) && s[idx+2] == '#' {
		tmp, err := strconv.Atoi(s[idx : idx+2])
		if err == nil && tmp > 9 && tmp <= 26 {
			// 因为题目说唯一解，所以可以直接返回
			return string(byte(tmp-1+'a')) + process(idx+3, s)
		}
	}
	return string(s[idx]-'1'+'a') + process(idx+1, s)
}

// @lc code=end
