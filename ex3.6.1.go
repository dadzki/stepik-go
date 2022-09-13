package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Result struct {
	Average float64
}

func main() {
	var group Group

	//if err := json.NewDecoder(os.Stdin).Decode(&group); err != nil {
	//	panic(err)
	//}
	//
	// OR:
	//
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	data = []byte(`
		{
			"ID":134,
			"Number":"ИЛМ-1274",
			"Year":2,
			"Students":[
				{
					"LastName":"Name1",
					"FirstName":"фамил1",
					"MiddleName":"отч1",
					"Birthday":"4апреля1970года",
					"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира21",
					"Phone":"+7(948)709-47-24",
					"Rating":[1,2,3]
				},
				{
					"LastName":"Name2",
					"FirstName":"фамил2",
					"MiddleName":"отч1",
					"Birthday":"4апреля1970года",
					"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
					"Phone":"+7(948)709-47-24",
					"Rating":[4,2,4]
				}
			]
		}
	`)
	//fmt.Println(string(data))

	if err := json.Unmarshal(data, &group); err != nil {
		panic(err)
	}

	var rating int
	for _, student := range group.Students {
		rating += len(student.Rating)
	}

	result := Result{
		Average: float64(rating) / float64(len(group.Students)),
	}
	dataJson, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataJson))
}
