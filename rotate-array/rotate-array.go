package main

import "fmt"

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		tmp := a[i]
		a[i] = a[n-1-i]
		a [n-1-i] = tmp
	}
}

func rotate(nums []int, k int) {
	k =  k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}


func main() {
	nums := []int{-4, -3, 0, 1, 2, 3, 4,6,7,8,9}
	rotate(nums, 3)
	fmt.Printf("ret:%v\n", nums)
}