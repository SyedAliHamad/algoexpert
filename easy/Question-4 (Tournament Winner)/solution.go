package main

import "fmt"

func FindWinner(competetion [][]string,result []int) string {

    return ""
}

func main() {
	var matches [][]string

    // Initialize the 2D slice
    matches = append(matches, []string{"apple", "orange", "banana"})
    matches = append(matches, []string{"cat", "dog", "rabbit"})
    matches = append(matches, []string{"red", "green", "blue"})

	var results = []int{0,1,0,1,0,1}

    winner:=FindWinner(matches, results)
    fmt.Sprintln("Winner of this tournament is ", winner)
}
