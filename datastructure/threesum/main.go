package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans = make([][]int, 0)
	sort.Ints(nums)
	arrLen := len(nums)
	var l1 = 0
	for i := 0; i < arrLen-2; i++ {
		for i !=0 && nums[i]== nums[l1] && i< arrLen-2{
			i++
		}
		l1 = i
		var l2,l3 = i+1, arrLen-1
		for l2 < l3 {
			if nums[l2] + nums[l3] > -nums[l1] {
				l3--
			}else if nums[l2] + nums[l3] < -nums[l1]{
				l2++
			}else{
				var temp = make([]int,3)
				temp[0], temp[1], temp[2] = nums[l1], nums[l2], nums[l3]
				ans  = append(ans, temp)
				for nums[l2] == nums[l2+1]&& l2 < l3 && l2 < arrLen-2{
					l2 ++
				}
				for nums[l3] == nums[l3-1]&& l2 < l3 && l3 > 1{
					l3--
				}
				l2++
				l3--
			}

		}
	}
	return ans
}
func main() {
   var nums = []int{-1,0,1,2,-1,-4}
   ans := threeSum(nums)
   fmt.Println(ans)
   nums = []int{0,0,0}
	ans = threeSum(nums)
	fmt.Println(ans)
}
