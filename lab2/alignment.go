package main

import (
	"fmt"
	"os"
)

const INFILE = "/Users/coordinate36/program/algorithm/lab2/alignment.txt"

var len1, len2 int
var str1, str2 []int
var gapPenality []int
var mismatchPenality [][]int

func input() {
	var m int
	fmt.Scanf("%d", &m)
	gapPenality = make([]int, m)
	mismatchPenality = make([][]int, m)
	for i := range mismatchPenality {
		mismatchPenality[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &gapPenality[i])
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &mismatchPenality[i][j])
		}
	}

	fmt.Scanf("%d", &len1)
	str1 = make([]int, len1)
	for i := 0; i < len1; i++ {
		fmt.Scanf("%d", &str1[i])
	}
	fmt.Scanf("%d", &len2)
	str2 = make([]int, len2)
	for i := 0; i < len2; i++ {
		fmt.Scanf("%d", &str2[i])
	}

	fmt.Scanln()
}

func min(first int, args ...int) int {
	for _, v := range args {
		if first > v {
			first = v
		}
	}
	return first
}

func align() int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len2; i++ {
		dp[0][i] = dp[0][i-1] + gapPenality[str2[i-1]]
	}
	for i := 1; i <= len1; i++ {
		dp[1][0] += gapPenality[str1[i-1]]
		for j := 1; j <= len2; j++ {
			dp[1][j] = min(dp[1][j-1]+gapPenality[str2[j-1]], dp[0][j-1]+mismatchPenality[str1[i-1]][str2[j-1]], dp[0][j]+gapPenality[str1[i-1]])
		}
		for j := 0; j <= len2; j++ {
			dp[0][j] = dp[1][j]
		}
	}

	return dp[1][len2]
}

func main() {

	// redirect stdin
	inFile, err := os.Open(INFILE)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()
	os.Stdin = inFile

	var numTests int
	fmt.Scanf("%d", &numTests)
	fmt.Scanln()

	for i := 0; i < numTests; i++ {
		input()
		minPenality := align()
		fmt.Println("minimum penality:", minPenality)
	}
}
