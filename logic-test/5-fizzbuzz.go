package logic_test

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) []string {
	arr := make([]string, n)
	fizz := "Fizz"
	buzz := "Buzz"
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			arr[i-1] = fizz + buzz
		} else if i%3 == 0 {
			arr[i-1] = fizz
		} else if i%5 == 0 {
			arr[i-1] = buzz
		} else {
			arr[i-1] = strconv.Itoa(i)
		}
	}
	fmt.Println(arr)
	return arr
}
