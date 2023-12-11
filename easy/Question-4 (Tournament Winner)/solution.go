package main

import (
	"fmt"
)

func FindWinner(competetion [][]string,result []int) string {

    totalMatches:=len(competetion)
    winCount:=make(map[string]int)

    for i:=0; i<totalMatches; i++ {
        competitor1:=competetion[i][0]
        competitor2:=competetion[i][1]
        if result[i]== 0 {
            val,exist:=winCount[competitor2]
            if exist{
                winCount[competitor2]=val+1
            } else{
                winCount[competitor2]=1
            }
        } else {
            val,exist:=winCount[competitor1]
            if exist{
                winCount[competitor1]=val+1
            } else{
                winCount[competitor1]=1
            }
        }
    }

    maxcount:=0
    Winner:=""
    for key, value := range winCount{
        if maxcount<value{
            maxcount=value
            Winner=key
        }
    }
    
    return Winner
}

func main() {
	var matches [][]string

    // Initialize the 2D slice
    matches = append(matches, []string{"html", "c++"})
    matches = append(matches, []string{"html", "python"})
    matches = append(matches, []string{"html", "go"})
    matches = append(matches, []string{"html", "java"})
    matches = append(matches, []string{"c++", "python"})
    matches = append(matches, []string{"c++", "go"})
    matches = append(matches, []string{"c++", "java"})
    matches = append(matches, []string{"python", "go"})
    matches = append(matches, []string{"python", "java"})
    matches = append(matches, []string{"go", "java"})

	var results = []int{0,0,0,0,0,0,1,0,1,1}

    winner:=FindWinner(matches, results)

    fmt.Println("Winner of this tournament is ", winner)
}
