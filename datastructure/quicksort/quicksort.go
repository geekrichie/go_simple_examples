package main

import (
	"fmt"
	"math/rand"
	"time"
)
func swap(num []int, i, j int) {
	num[i], num[j]= num[j], num[i]
}

func quickSort(num []int, start, end int) {
     if start >= end {
     	return
	 }
	 var idx = rand.Int()%(end+1-start)+start
	 swap(num,start,idx)
	 var j  = start+1
	 for i := start+1; i <= end; i++ {
	 	if num[i] < num[start] {
	 		swap(num, i, j)
	 		j++
		}
	 }
	 swap(num,start,j-1)
	 quickSort(num, start, j-2)
	 quickSort(num, j,end)
}

func main() {
	rand.NewSource(time.Now().Unix())
	var b  = []int {10,6,8,12,1,87,34,2}
	//var b = []int{1,0}
	quickSort(b, 0, len(b)-1)
	fmt.Println(b)
}