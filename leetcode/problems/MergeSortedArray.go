// https://leetcode.com/problems/merge-sorted-array/
package problems

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = nums2
	}
	var n1Idx, n2Idx int

	// Merge until one list runs out of elements
	for n1Idx < m && n2Idx < n {
		if nums1[n1Idx] < nums2[n2Idx] {
			n1Idx += 1
		} else {
			temp := append([]int{nums2[n2Idx]}, nums1[n1Idx:m+n-1]...)
			nums1 = append(nums1[:n1Idx], temp...)
			n1Idx += 1
			n2Idx += 1
		}
	}

	// nums2 may have elements larger than anything in nums1.
	// Make sure to append the last elements of nums2 if this is true
	if n2Idx < n {
		nums1 = append(nums1[:n1Idx+1], nums2[n2Idx:]...)
	}
}

func mergeWithSort(nums1 []int, m int, nums2 []int, n int) {
	// Borrowed from leetcode solutions
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}

func MergeSortedArray() {
	var (
		nums1 []int
		m     int // Number of elements in nums1
		nums2 []int
		n     int //Number of elements in nums2
	)
	fmt.Println("Running LeetCode problem: Merge Sorted Arrays")
	fmt.Println("Test case 1: ")
	m = 3
	n = 3
	nums1 = append([]int{1, 2, 3}, make([]int, n)...)
	nums2 = []int{2, 5, 6}
	// merge(nums1, m, nums2, n)
	mergeWithSort(nums1, m, nums2, n)
	fmt.Printf("Answer: %v\n", nums1)
}
