/* a simiple program to input numbers and give output
the sum of numbers*/
package main

import (
	"fmt"
)

func main() {
	var number int //lenght of the array*
	var temp int   // variable into which the numbers from stdin would be stored
	sum := 0       //final sum to be output to stdout
	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		fmt.Scan(&temp)
		sum = sum + temp
	}
	fmt.Print(sum)
}
