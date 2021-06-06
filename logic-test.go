package main

import (
	"time"

	logic_test "github.com/raismaulana/rid-apply/logic-test"
)

func main() {
	logic_test.Palindrome("Radar")
	logic_test.Leap_year()
	time.Sleep(1 * time.Second)
	logic_test.Reverse()
	logic_test.Nearest_fib([]int{15, 1, 3})
	logic_test.FizzBuzz(15)
}
