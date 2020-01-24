package main

/*
 * @lc app=leetcode.cn id=1318 lang=golang
 *
 * [1318] 或运算的最小翻转次数
 */

// @lc code=start
func minFlips(a int, b int, c int) int {

	t := a | b
	x := t ^ c
	cnt := 0
	for i := 0; i < 32; i++ {

		x1 := x >> uint(i)
		a1 := a >> uint(i)
		b1 := b >> uint(i)
		c1 := c >> uint(i)
		if 1&x1 == 1 {
			// x右移i位，跟1做与操作可以得到x的第i位是0还是1
			cBit := 1 & c1
			if cBit == 0 {

				aBit := 1 & a1
				bBit := 1 & b1

				if aBit == 1 && bBit == 1 {
					cnt += 2
				} else {
					cnt++
				}

			} else {
				cnt++
			}
		}

	}
	return cnt
}

// @lc code=end

//func main() {
//	c := minFlips(8, 3, 5)
//	fmt.Println(c)
//}
