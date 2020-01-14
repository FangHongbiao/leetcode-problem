package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=1320 lang=golang
 *
 * [1320] 二指输入的的最小距离
 */

// @lc code=start

const (
	BaseChar byte = 'A'
	N             = 27
)

// GetPosByChar得到给定字符在键盘上的左边
func GetPosByChar(c byte) (int, int) {

	offset := int(c) - int(BaseChar)
	return offset / 6, offset % 6
}

// 计算c1、c2之间的距离，注意，当其中存在开始字符时，说明手指处于起始状态，无代价
func Distance(c1, c2 byte) int {

	if c1 == '[' || c2 == '[' {
		return 0
	}

	x1, y1 := GetPosByChar(c1)
	x2, y2 := GetPosByChar(c2)
	return (int)(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func minimumDistance(word string) int {

	return dfs('[', '[', 0, word)
}

// dfs执行搜索过程,含义为：
// 左手在字符cl上，右手在字符cr上，搞定从idx到字符串word结尾，所需要的cost
var cache map[string]int = make(map[string]int, 0)

func dfs(cl byte, cr byte, idx int, word string) int {

	if idx == len(word) {
		return 0
	}
	key := fmt.Sprintf("%v-%v-%v", cl, cr, idx)
	if val, ok := cache[key]; ok {
		return val
	}

	// 用左手去按idx位置上的字符
	costL := Distance(cl, word[idx]) + dfs(word[idx], cr, idx+1, word)

	// 用左手去按idx位置上的字符
	costR := Distance(cr, word[idx]) + dfs(cl, word[idx], idx+1, word)

	var res int
	if costL < costR {
		res = costL
	} else {
		res = costR
	}

	cache[key] = res
	return res
}

func minimumDistance2(word string) int {
	word = word + " "
	dp := make([][][]int, N)
	// 初始化dp数组
	for i := 0; i < N; i++ {
		dp[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]int, len(word))
		}
	}
	// 填充dp数组
	fillDP(word, dp)
	return dp[N-1][N-1][0]
}

func fillDP(word string, dp [][][]int) [][][]int {
	// 由于N为27， 加了1，所以不需要额外处理最后一行，最后一行都是0，这个是正确的。
	// 因为此时最后一行的含义是，没有看见字符，所以左右手在任何位置的代价都可以视为0
	for L := len(word) - 2; L >= 0; L-- {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				costL := dp[int(word[L]-'A')][j][L+1] + Distance(word[L], byte(int('A')+i))
				costR := dp[i][int(word[L]-'A')][L+1] + Distance(byte(int('A')+j), word[L])
				if costL < costR {
					dp[i][j][L] = costL
				} else {
					dp[i][j][L] = costR
				}
			}
		}
	}
	return dp
}

func main() {
	d := minimumDistance2("NEW")
	fmt.Println(d)
}
