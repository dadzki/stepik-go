package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	num := 0
	for in.Scan() {
		if len(in.Text()) == 0 {
			break
		}
		n, err := strconv.Atoi(in.Text())
		if err != nil {
			panic(err)
		}

		num += n
	}

	_, _ = os.Stdout.WriteString(strconv.Itoa(num))
	// OR:
	//_, _ = io.WriteString(os.Stdout, strconv.Itoa(num))
	//
	// OR:
	//out := bufio.NewWriter(os.Stdout)
	//_, _ = out.WriteString(strconv.Itoa(num))
	//_ = out.Flush()
}
