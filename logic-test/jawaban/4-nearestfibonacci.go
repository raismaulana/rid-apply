package logic_test

import "fmt"

func Nearest_fib(arr []int) {
	// sum array
	var sum int
	for _, v := range arr {
		sum = sum + v
	}
	if sum < 0 {
		fmt.Println("Err: array sum must be a positif int")
		return
	}
	// searching the nearest fib gap
	fibNumBefore := 0 // var for store fib number before sum
	output := 0       // gap
	a, b := 0, 1      // fib starting num
	for {
		a, b = b, a+b
		if a > sum {
			x := sum - fibNumBefore
			y := a - sum
			if x < y {
				output = -x
			} else {
				output = y
			}
			break
		}
		fibNumBefore = a
	}
	fmt.Println(output)
}
