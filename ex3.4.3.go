package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	charge int
}

func (b Battery) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", b.charge))
}

func main() {
	var chargeStr string

	fmt.Scan(&chargeStr)

	batteryForTest := Battery{
		charge: strings.Count(chargeStr, "1"),
	}

	fmt.Println(batteryForTest)

}
