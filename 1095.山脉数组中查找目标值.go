package main

/*
 * @lc app=leetcode.cn id=1095 lang=golang
 *
 * [1095] 山脉数组中查找目标值
 */

// @lc code=start
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

type MountainArray struct {
	arr []int
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}
func (this *MountainArray) length() int {
	return len(this.arr)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	return help(target, 0, mountainArr.length()-1, mountainArr)
}

// 在山脉数组mountainArr的[L, R]区间里搜索目标target, 返回最小下标
func help(target int, L int, R int, mountainArr *MountainArray) int {

	if L == R {
		if mountainArr.get(L) == target {
			return L
		}
		return -1
	}

	mid := L + (R-L)/2

	sL := -1
	// 如果单调直接二分搜索，否则递归该过程
	if isMonotonous(mountainArr, L, mid) {
		sL = binarySearch(mountainArr, L, mid, target, true)
	} else {
		sL = help(target, L, mid, mountainArr)
	}

	// 如果左侧找到了，就直接结束，因为找的是最小的
	if sL != -1 {
		return sL
	}

	sR := -1
	if isMonotonous(mountainArr, mid+1, R) {
		sR = binarySearch(mountainArr, mid+1, R, target, false)
	} else {
		sR = help(target, mid+1, R, mountainArr)
	}

	return sR
}

// 最难的一个函数，如何判断[L,R]上是否单调，思考了很久
func isMonotonous(mountainArr *MountainArray, L int, R int) bool {

	vL := mountainArr.get(L)
	vR := mountainArr.get(R)

	// 因为能进入这个函数，至少L<R，且都在合法区间内，所以R-1,L+1索引一定是合法的
	if vL < vR && (vR > mountainArr.get(R-1)) ||
		vL > vR && (mountainArr.get(L+1) < vL) {
		return true
	}

	return false
}

func binarySearch(mountainArr *MountainArray, L int, R int, target int, asc bool) int {

	if L > R {
		return -1
	}

	if L == R {
		if mountainArr.get(L) == target {
			return L
		}
		return -1
	}

	mid := L + (R-L)/2

	vMid := mountainArr.get(mid)

	if vMid == target {
		return mid
	}

	if asc {
		if target < vMid {
			return binarySearch(mountainArr, L, mid-1, target, asc)
		} else {
			return binarySearch(mountainArr, mid+1, R, target, asc)
		}
	} else {
		if target > vMid {
			return binarySearch(mountainArr, L, mid-1, target, asc)
		} else {
			return binarySearch(mountainArr, mid+1, R, target, asc)
		}
	}

}

// @lc code=end

//func main() {
//	arr := &MountainArray{
//		arr: []int{
//			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82,
//		},
//	}
//	res := findInMountainArray(81, arr)
//	fmt.Println(res)
//}
