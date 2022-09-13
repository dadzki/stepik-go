package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n uint
	fmt.Scan(&n)

	fn := func(n uint) uint {
		var s string
		bs := []byte(strconv.FormatInt(int64(n), 10))

		for _, v := range bs {
			if v%2 == 0 && string(v) != "0" {
				s += string(v)
			}
		}

		num, _ := strconv.Atoi(s)
		if num == 0 {
			num = 100
		}
		return uint(num)
	}

	fmt.Println(fn(n))

}
