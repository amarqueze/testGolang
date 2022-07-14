package main

import (
	"fmt"
	"math"
	"time"
)

func plusPlus(a, b, c int) int {
	return a + b + c
}

func plusPlusUltra(initial int) func(a, b int) int {
	return func(a, b int) int {
		return plusPlus(initial, a, b)
	}
}

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	callback := plusPlusUltra(10)
	fmt.Println(callback(5, 5))
	fmt.Println(callback(15, 5))

	times()
}

func values() {
	fmt.Println("Hello", "Word", "Golang")

	fmt.Println("1+2 =", 2+1)
	fmt.Println("8.0/3.0 =", 8.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variables() {
	const PI = 3.1415

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 8
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "Batman"
	fmt.Println(f)

	fmt.Println(math.Sin(PI))

	var h = float64(PI)
	var v = int(h)
	fmt.Println(v)
}

func control() {
	if 1 == 1 {
		fmt.Println("is Equals")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

func cycles() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	arrayNames := []string{"Alan", "Berton"}
	for k, v := range arrayNames {
		fmt.Println("Key: ", k, "Value: ", v)
	}

	for _, v := range arrayNames {
		fmt.Println("value: ", v)
	}
}

func maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
}

func arrays() {
	var a [5]int
	fmt.Println("array:", a)

	s := make([]string, 3)
	fmt.Println("slice:", s)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func times() {
	now := time.Now()
	fmt.Println(now.Weekday().String())
}
