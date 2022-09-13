package main

import (
	"fmt"
)

func main() {

	groupCity := map[int][]string{
		10:   {"город1", "город2", "город3"},
		100:  {"город4", "город5", "город6"},
		1000: {"город7", "город8", "город9"},
	}
	fmt.Println(groupCity)

	cityPopulation := map[string]int{
		"город1": 10,
		"город2": 20,
		"город3": 30,
		"город4": 100,
		"город5": 200,
		"город6": 300,
		"город7": 1000,
		"город8": 2000,
		"город9": 3000,
	}

	for number, cities := range groupCity {
		if number == 100 {
			continue
		}

		for _, city := range cities {
			delete(cityPopulation, city)
		}
	}
	fmt.Println(cityPopulation)
}
