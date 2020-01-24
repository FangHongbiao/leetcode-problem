/*
 * @lc app=leetcode.cn id=1290 lang=golang
 *
 * [1290] 二进制链表转整数
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans = head.Val + (ans << 1)
		head = head.Next
	}
	return ans
}

// @lc code=end
