package main

import "fmt"

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	if (numsLen == 0) || target < nums[0] {
		return 0
	}
	if  target > nums[numsLen-1]{
		return numsLen
	}
	leftIndex := 0
	rightIndex := numsLen-1
	for {
		midIndex := (leftIndex + rightIndex)/2
		fmt.Printf("left[%d]:%d, right[%d]:%d, mid[%d]:%d\n", leftIndex, nums[leftIndex], rightIndex, nums[rightIndex], midIndex, nums[midIndex])
		if target == nums[midIndex]{
			return midIndex
		}else if target < nums[midIndex]{
			rightIndex = midIndex
		}else{
			leftIndex = midIndex
		}
		if rightIndex - leftIndex <= 1{
			if target == nums[rightIndex] {
				return rightIndex
			}else if target == nums[leftIndex]{
				return leftIndex
			}else{
				return rightIndex
			}
		}
	}

}

func main() {
	nums := []int{1, 2, 3, 4,6,7,8,9}
	index := searchInsert(nums, 5)
	fmt.Printf("index:%d\n", index)
}