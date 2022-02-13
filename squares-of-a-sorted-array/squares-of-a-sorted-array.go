package main

import "fmt"

func sortedSquares(nums []int) []int {
	numsLen := len(nums)
	left := 0
	right := numsLen -1
	newNums  := make([]int, numsLen)
	i := numsLen-1
	for{
		squre_left := nums[left]*nums[left]
		squre_right := nums[right]*nums[right]
		if right == left{
			newNums[i] = squre_left
			return newNums
		}
		if squre_left > squre_right{
			newNums[i] = squre_left
			left = left + 1
		}else{
			newNums[i] = squre_right
			right = right -1
		}
		i = i-1
	}
}

func main() {
	nums := []int{-4, -3, 0, 1, 2, 3, 4,6,7,8,9}
	ret := sortedSquares(nums)
	fmt.Printf("ret:%v\n", ret)
}