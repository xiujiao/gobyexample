package main

import (
	"fmt"
	"math"
	"time"
)

func add(a int, b int) int {
	return a + b
}

func sum(num ...int) int {
	total := 0
	for _, value := range num {
		total = total + value
	}
	return total
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func value(n int) {
	n = 0
}

func pointer(n *int) {
	*n = 0
}

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

	s := make([]string, 3)
	s[0] = "x"
	s[1] = "i"
	s[2] = "u"
	fmt.Println("s is ", s, "the length of it is ", len(s))

	s = append(s, "jiao")
	s = append(s, "is")
	s = append(s, "wonderful")

	fmt.Println("now s is :", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c is ", c)

	l := s[2:5]
	l1 := s[:5]
	l2 := s[2:]
	fmt.Println(l, l1, l2)

	twoD := make([][]int, 2)
	for i := 0; i < 2; i++ {
		innerlen := i + 3
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	for i, content := range s {
		fmt.Println(i, content)
	}

	ans := add(1, 3)
	fmt.Println("1+3=", ans)

	num := []int{2, 4, 4, 6, 7, 0, 9}
	ans = sum(num...)
	fmt.Println("the sum is ", ans)

	fmt.Println(" the factorial  of 7 is ", fact(7))

	i := 1

	fmt.Println("i is ", i)
	value(i)
	fmt.Println("after calling value func, i is ", i)
	pointer(&i)
	fmt.Println("after calling pointer funcm i is ", i)

}
