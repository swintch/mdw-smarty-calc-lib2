package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

func main() {
	calculator := &calc.Addition{}
	calcValue1 := ParsInt(os.Args[1])
	calcValue2 := ParsInt(os.Args[2])
	_, err := fmt.Println(calculator.Calculate(calcValue1, calcValue2))
	if err != nil {
		panic(err)
	}
}

func ParsInt(stringValue string) int {
	integerValue, err := strconv.Atoi(stringValue)
	if err != nil {
		panic(err)
	}
	return integerValue
}
