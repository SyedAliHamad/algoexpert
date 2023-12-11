package main

import (
	"fmt"
	"sort"
)

func LowestNonExistantChange(coins []int) int {
	sort.Ints(coins)
	hPC := 0
	for _,v:=range coins {
		if v>hPC+1{
			return hPC+1
		} else {
		hPC=hPC+v
		}
	}

	return 0
}
func main() {
	coins := []int{5, 7, 1, 1, 2, 3, 22}
	L := LowestNonExistantChange(coins)

	fmt.Println("lowest change not possible is:", L )
}