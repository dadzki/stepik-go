package main

import "fmt"

func main() {
	// a, b := sumInt(1, 0)
	// fmt.Println(a, b)

	x1 := 2
	x2 := 4
	test(&x1, &x2)

}

func sumInt(arg ...int) (count int, summ int) {
	for _, v := range arg {
		summ += v
		count++
	}
	return
}

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}
