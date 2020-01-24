package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1235 lang=golang
 *
 * [1235] 规划兼职工作
 */

// @lc code=start

// Work 用于保存工作信息，包括: 开始时间、结束时间、利润
type Work struct {
	startTime int
	endTime   int
	profit    int
	index     int
}

// golang 如何使用堆
type PriorityQueue []*Work

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].profit/(pq[i].endTime-pq[i].startTime) > pq[j].profit/(pq[j].endTime-pq[j].startTime)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Work)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func jobScheduling1(startTime []int, endTime []int, profit []int) int {

	pq := make(PriorityQueue, 0, len(startTime))

	for i := 0; i < len(startTime); i++ {
		pq.Push(&Work{
			startTime: startTime[i],
			endTime:   endTime[i],
			profit:    profit[i],
		})
	}
	heap.Init(&pq)

	totalProfit := 0
	cands := make([]*Work, 0)
	for pq.Len() > 0 {
		cand := heap.Pop(&pq).(*Work)
		fmt.Println(*cand)
		if !overlap(cands, cand) {
			totalProfit += cand.profit
			cands = append(cands, cand)
		}
	}

	return totalProfit
}

func overlap(works []*Work, work *Work) bool {

	for _, w := range works {
		// 重叠的两种情况: w的开始或结束时间在work的开始或结束时间内部 || work的开始或结束时间在w的开始或结束时间内部
		if work.startTime > w.startTime && work.startTime < w.endTime ||
			work.endTime > w.startTime && work.endTime < w.endTime ||
			w.startTime > work.startTime && w.startTime < work.endTime ||
			w.endTime > work.startTime && w.endTime < work.endTime {
			return true
		}
	}

	return false
}

func jobScheduling2(startTime []int, endTime []int, profit []int) int {

	endsMap := make(map[int][]int, 0)
	maxEndTime := 0
	for i := 0; i < len(startTime); i++ {
		if _, ok := endsMap[endTime[i]]; !ok {
			endsMap[endTime[i]] = make([]int, 0)
		}
		endsMap[endTime[i]] = append(endsMap[endTime[i]], i)

		// 更新最大结束时间
		if endTime[i] > maxEndTime {
			maxEndTime = endTime[i]
		}
	}

	dp := make([]int, maxEndTime+1)

	for i := 1; i <= maxEndTime; i++ {
		dp[i] = dp[i-1]
		if _, ok := endsMap[i]; ok {
			for _, prev := range endsMap[i] {
				if dp[startTime[prev]]+profit[prev] > dp[i] {
					dp[i] = dp[startTime[prev]] + profit[prev]
				}
			}
		}
	}

	return dp[maxEndTime]
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][]int, 0)
	for i := 0; i < n; i++ {
		jobs = append(jobs, []int{startTime[i], endTime[i], profit[i]})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})
	dp := make([]int, n)
	dp[0] = jobs[0][2]

	for i := 1; i < n; i++ {
		l, r := 0, i-1
		for l < r {
			mid := (l + r + 1) >> 1
			if jobs[mid][1] <= jobs[i][0] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if jobs[l][1] <= jobs[i][0] {
			dp[i] = max(dp[i-1], dp[l]+jobs[i][2])
		} else {
			dp[i] = max(dp[i-1], jobs[i][2])
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
