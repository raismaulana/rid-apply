package main

import (
	"time"

	logic_test "github.com/raismaulana/rid-apply/logic-test/jawaban"
)

func main() {
	logic_test.Palindrome("Radar")
	time.Sleep(1 * time.Second)
	logic_test.Leap_year()
	time.Sleep(1 * time.Second)
	logic_test.Reverse()
	time.Sleep(1 * time.Second)
	logic_test.Nearest_fib([]int{15, 1, 3})
	time.Sleep(1 * time.Second)
	logic_test.FizzBuzz(15)
}
