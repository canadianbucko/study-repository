package main

import (
	"fmt"
	"loh/calc"
)

func main() {
	c := calc.New()
	result, err := c.Sum(12, 20)
	fmt.Println(result, err)
	resultDivide, err := c.Divide(12, 10)
	fmt.Println(resultDivide, err)

}
