package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1249 lang=golang
 *
 * [1249] 移除无效的括号
 */

// @lc code=start

var longest string

func minRemoveToMakeValid1(s string) string {
	// 初始化全局最长的(最长的说明移除的最少)
	longest = ""
	// 当前留下来的字符串
	cur := ""
	dfs(s, 0, 0, 0, cur)
	return longest
}

func dfs(s string, idx int, left int, right int, cur string) {

	if right > left {
		return
	}

	if idx == len(s) {
		fmt.Println(cur)
		if len(cur) > len(longest) && left == right {
			longest = cur
		}
		return
	}

	if s[idx] == ')' {
		// 移除
		dfs(s, idx+1, left, right, cur)
		// 不移除
		dfs(s, idx+1, left, right+1, cur+(string)(s[idx]))
	} else if s[idx] == '(' {
		// 移除
		dfs(s, idx+1, left, right, cur)
		// 不移除
		dfs(s, idx+1, left+1, right, cur+(string)(s[idx]))
	} else {
		dfs(s, idx+1, left, right, cur+(string)(s[idx]))
	}
}

func minRemoveToMakeValid23(s string) string {
	return dfs2(s, 0, 0, 0, "")
}

func dfs2(s string, idx int, left int, right int, cur string) string {

	if right > left {
		return ""
	}

	if idx == len(s) {
		if left == right {
			return cur
		}
		return ""
	}

	if s[idx] == ')' {
		// 移除
		removed := dfs2(s, idx+1, left, right, cur)
		// 不移除
		noRemoved := dfs2(s, idx+1, left, right+1, cur+(string)(s[idx]))

		return min(removed, noRemoved)

	} else if s[idx] == '(' {
		// 移除
		removed := dfs2(s, idx+1, left, right, cur)
		// 不移除
		noRemoved := dfs2(s, idx+1, left+1, right, cur+(string)(s[idx]))

		return min(removed, noRemoved)
	} else {
		return dfs2(s, idx+1, left, right, cur+(string)(s[idx]))
	}
}

func min(i, j string) string {
	// 返回空串说明无效, 直接取另一个即可
	if i == "" {
		return j
	}

	if j == "" {
		return i
	}

	// 非空串, 返回更长的, 意味着删除的少
	if len(i) > len(j) {
		return i
	}
	return j
}

var cache map[string]string

func minRemoveToMakeValid3(s string) string {
	cache = make(map[string]string)
	return dfs3(s, 0, 0, 0, "")
}

func dfs3(s string, idx int, left int, right int, cur string) string {

	if right > left {
		return ""
	}

	if idx == len(s) {
		if left == right {
			return cur
		}
		return ""
	}

	key := fmt.Sprintf("%d_%d_%d", idx, left, right)

	if value, ok := cache[key]; ok {
		return value
	}

	if s[idx] == ')' {
		// 移除
		removed := dfs3(s, idx+1, left, right, cur)
		// 不移除
		noRemoved := dfs3(s, idx+1, left, right+1, cur+(string)(s[idx]))

		cache[key] = min(removed, noRemoved)

	} else if s[idx] == '(' {
		// 移除
		removed := dfs3(s, idx+1, left, right, cur)
		// 不移除
		noRemoved := dfs3(s, idx+1, left+1, right, cur+(string)(s[idx]))

		cache[key] = min(removed, noRemoved)
	} else {
		cache[key] = dfs3(s, idx+1, left, right, cur+(string)(s[idx]))
	}
	return cache[key]
}

func minRemoveToMakeValid(s string) string {
	// string 类型不可变, 先转换成[]byte
	b := []byte(s)
	idxL := []int{}
	left, right := 0, 0
	for i := 0; i < len(b); i++ {
		if s[i] == '(' {
			idxL = append(idxL, i)
			left++
		} else if b[i] == ')' {
			if right >= left {
				b[i] = 0
			} else {
				right++
			}
		}
	}

	for i := 0; i < left-right; i++ {
		// 要从后往前移除多余的'(', 因为前面的可能已经配对了
		b[idxL[len(idxL)-1-i]] = 0
	}

	// 收集结果
	ans := []byte{}
	for _, c := range b {
		if c != 0 {
			ans = append(ans, c)
		}
	}

	return string(ans)
}

// @lc code=end
