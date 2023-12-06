package main

import "log"

func sorted_square(arr []int) []int {

	new_arr := make([]int, len(arr))

	left := 0
	right := len(arr) - 1
	log.Println(left, " ", right)
	for i := 0; i < len(arr); i++ {
		if arr[left] > arr[right] {
			num := arr[left]
			new_arr[len(arr)-i-1] = num * num
			left += 1
		} else {
			num := arr[right]
			new_arr[len(arr)-i-1] = num * num
			right -= 1
		}
		log.Println(new_arr)
	}

	return new_arr
}

func main() {
	arr := []int{-7, -4, -3, -1, 1, 5, 6, 8, 9}
	new_arr := sorted_square(arr)

	log.Println(new_arr)
}
