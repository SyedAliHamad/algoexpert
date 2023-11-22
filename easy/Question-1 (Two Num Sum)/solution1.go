package main

import (
	"fmt"
	"sort"
)

func twoNumSum(arr []int, targetSum int) (int,int) {
	sort.Ints(arr)
	max:=len(arr)-1
	L := 0
	R := max

	for i:=0;i<=max;i++ {
		sum:=arr[L]+arr[R]
		if sum == targetSum {
			return arr[L], arr[R]
		}
		if sum < targetSum {
			L=L+1
		} else if sum > targetSum{
			R=R-1
		}
	}

	return 0,0
}

func main() {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	num := 23

	fmt.Println(twoNumSum(arr, num))
}
