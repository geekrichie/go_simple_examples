package main

import (
	"fmt"
)

var MaxInt = 1 << 31 - 1

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0{
		return 0
	}
	var m,n =len(grid)+1, len(grid[0])+1
	var dp = make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0 ; i< m; i++ {
		dp[i][0] = MaxInt
	}
	for i := 0 ; i< n; i++ {
		dp[0][i] = MaxInt
	}
	for i := 1; i < m; i ++ {
		for j := 1; j< n ; j++ {
			if i == 1 && j == 1 {
				dp[i][j] = grid[i-1][j-1]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Println()
	}
	return dp[m-1][n-1];
}

func min (x ,y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	var square = [][]int{
		[]int{1,3,1},
		[]int{1,5,1},
		[]int{4,2,1},
	}
	ans := minPathSum(square)
	fmt.Println(ans)
}