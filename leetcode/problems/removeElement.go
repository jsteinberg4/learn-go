// Package problems https://leetcode.com/problems/remove-element/
package problems

import (
	"log/slog"
)

func RemoveElement() {
	// case1()
	// case2()
	case3()
}

// swap Swaps arr[i] with arr[j] in place
func swap[T any](arr []T, i, j int) {
	// WARNING: Will break if i,j out of range
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func removeElement(nums []int, val int) int {
	// Tasks:
	// 1) Remove occurences of val in place
	// 2) Count # elements != val
	var (
		countVals int // k
		left      int = 0
		right     int = len(nums) - 1
	)

	// Make sure right doesn't point to `val`
	for right > 0 && nums[right] == val {
		countVals++
		right--
	}
	slog.LogAttrs(
		nil, slog.LevelDebug, "After decr right loop",
		slog.Group("args", slog.Int("val", val), slog.Any("nums", nums)),
		slog.Group("locals",
			slog.Int("countVals", countVals),
			slog.Int("left", left),
			slog.Int("right", right)),
	)

	// Iterate left->right. Swap `val` to the end and count as they're found
	// Let left<- index counting up, righ<- index counting down
	// Until left & right meet/swap
	//.  -> if left != val, continue
	//.  -> else swap left, right
	for left < right {
		if nums[left] == val {
			swap(nums, left, right)
			countVals++
			right--
		}
		left++
	}

	return len(nums) - countVals
}

func case1() {
	nums := []int{3, 2, 2, 3}
	val := 3
	_ = removeElement(nums, val)
}
func case2() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	logInputs("Case 2", nums, val)
	counts := removeElement(nums, val)
	logOutputs("Case 2", nums, counts)
}

func logInputs(msg string, nums []int, val int) {
	slog.LogAttrs(nil, slog.LevelInfo, msg, slog.Group(
		"inputs", slog.Int("val", val), slog.Any("nums", nums)),
	)
}
func logOutputs(msg string, nums []int, counts int) {
	slog.LogAttrs(nil, slog.LevelInfo, msg,
		slog.Group("results", slog.Int("valueCounts", counts), slog.Any("nums", nums)),
	)
}

func case3() {
	nums := []int{2, 1, 2, 0, 2, 6}
	val := 2
	logInputs("Case 3", nums, val)
	counts := removeElement(nums, val)
	logOutputs("Case 3", nums, counts)
}
