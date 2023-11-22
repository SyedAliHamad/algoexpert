package main

import "log"

func sub_seq_exist(sequence []int, sub_seq []int) bool {
	sub_index:=0
	for i:=0;i<len(sequence);i++{
		if sequence[i]==sub_seq[sub_index] {
			sub_index+=1
			if sub_index==len(sub_seq) {
				return true
			}
		}
	}

	return false
}

func main() {
	sequence := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sub_seq := []int{1, 5, 10}
	
	log.Println(sub_seq_exist(sequence, sub_seq))
}