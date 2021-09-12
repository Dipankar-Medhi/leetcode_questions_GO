package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}

	targetInNum := checkTargetInNum(nums, target)

	if !targetInNum {
		return ans
	}
	start := search(nums, target, true)
	end := search(nums, target, false)
	ans = []int{start, end}
	return ans
}

func search(nums []int, target int, findStart bool) int {
	pot_ans := -1
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			pot_ans = mid
			if findStart {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return pot_ans
}

func checkTargetInNum(nums []int, target int) bool {
	for _, b := range nums {
		if target == b {
			return true
		}
	}
	return false
}
