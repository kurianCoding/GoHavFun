// this is a program which takes in two arrays
// and gives as output the comparison of the two
// arrays
package main

import "fmt"

func main() {
	var temp int
	scoreA := make([]int, 3)
	sumA := 0
	sumB := 0
	for j := 0; j < 2; j++ {
		for i := 0; i < 3; i++ {
			if j == 0 {
				fmt.Scan(&scoreA[i])
			} else {
				fmt.Scan(&temp)
				scoreA[i] = compare(scoreA[i], temp)
			}

		}
	}
	for k := 0; k < 3; k++ {
		if scoreA[k] > 0 {
			sumA = sumA + scoreA[k]
		} else {
			sumB = sumB + -1*scoreA[k]
		}
	}
	fmt.Printf("%d %d", sumA, sumB)
}

// compare between two integers and check if one
// is greater than the other
func compare(number int, temp int) int {
	if temp == number {
		return 0
	}
	if temp < number {
		return 1
	}
	return -1
}
