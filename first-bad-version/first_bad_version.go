package main

func firstBadVersion(n int) int {
	left := 1
	right := n
	for {
		mid := (left+right)/2
		if isBadVersion(mid) {
			right = mid
		}else{
			left = mid+1
		}
		if right - left <= 1{
			if isBadVersion(left){
				return left
			}else{
				return right
			}
		}
	}
}