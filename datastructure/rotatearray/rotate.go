package main

import (
	"fmt"
)

func binarySearch(b []int, l, r, target int) bool{
	if target < b[l] || target > b[r] {
		return false
	}
	for l <= r {
		var mid =  (l+r)/2
		if b[mid] == target {
			return true
		}else if b[mid] > target {
			r = mid-1
		}else {
			l = mid+1
		}
	}
	return false
}

func findElement(b []int, target int) bool {
	var l ,r=0,len(b)-1
	for l <= r {
		var mid = (l+r)/2
		if b[mid] < b[r] {
			var bsResult = binarySearch(b, mid,r, target)
			if bsResult == true{
				return true
			}
			r = mid-1
		}else if   b[mid] == b[r]{
			if b[mid] == target {
				return true
			}
			r--
		}else {
			var bsResult = binarySearch(b, l,mid, target)
			if bsResult == true{
				return true
			}
			l = mid + 1
		}
	}
	return false
}

func main() {
	var b = []int{1,2,1}
	var ans = findElement(b, 2)
	fmt.Println(ans)
}
