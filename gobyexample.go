package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("hello, xiujiao")

	fmt.Println("go" + "long" + "a cool language")
	fmt.Println("1+1=", 1+1)
	fmt.Println("8.0/2.5=", 8.0/2.5)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	var a string = "love is luxury, not everyone will own it forever"
	var f = 0
	xiujiao := "xj"
	var past = 3.14
	var o, p, q = 2, 4, 7.89
	fmt.Println("a=", a)
	fmt.Println("f=", f)
	fmt.Println("xiujiao is ", xiujiao)
	fmt.Println("past is", past)
	fmt.Println("o,p.q", o, p, q)

	const today = "Saturday"
	const n = 2000
	fmt.Println(today)
	fmt.Println(math.Sin(n))

	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
	if j := 0; j < 6 {
		j = j + 3
	} else if j > 10 {
		j = j - 3
	}

	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("This is weekend")
	default:
		fmt.Println("this is a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("this is before noon")
	default:
		fmt.Println("this is afternoon")
	}
}
