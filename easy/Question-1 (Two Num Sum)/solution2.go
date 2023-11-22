package main

import (
	"fmt"
)

func twoNumSum(arr []int, targetSum int) (int,int) {
	hashtable:=make(map[int]int)

	for i:=0;i<len(arr);i++{
		x:=arr[i]
		y:=targetSum-x
		_,exists:=hashtable[y]
		if exists{
			return x,y
		}
		hashtable[x]=0
	}

	return 0,0
}

func main() {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	num := 23

	fmt.Println(twoNumSum(arr, num))
}
