package main

import (
	"errors"
	"fmt"
)

func main() {
	var fnOperation func(a float64, b float64) float64

	value1, value2, oper := readTask()

	number1, err := getValue(value1)
	if err != nil {
		fmt.Println(err)
		return
	}

	number2, err := getValue(value2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fnOperation, err = getOperation(oper)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%.4f\n", fnOperation(number1, number2))
}

func getOperation(val interface{}) (fn func(n1 float64, n2 float64) float64, err error) {
	v, ok := val.(string)

	if !ok {
		return nil, errors.New(fmt.Sprintf("не верный тип '%T', должен быть string", v))
	}

	switch v {
	case "+":
		fn = func(n1 float64, n2 float64) float64 {
			return n1 + n2
		}
	case "-":
		fn = func(n1 float64, n2 float64) float64 {
			return n1 - n2
		}
	case "*":
		fn = func(n1 float64, n2 float64) float64 {
			return n1 * n2
		}
	case "/":
		fn = func(n1 float64, n2 float64) float64 {
			return n1 / n2
		}
	default:
		err = errors.New("неизвестная операция")
	}

	return
}

func getValue(value interface{}) (float64, error) {
	if v, ok := value.(float64); ok {
		return v, nil
	}

	return 0, errors.New(fmt.Sprintf("value=%v: %T", value, value))
}

func readTask() (interface{}, interface{}, interface{}) {
	return 10.45, 1.45, "+"
}
