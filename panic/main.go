package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("there was a panic!", p)
		}
	}()
	/* a := 3
	b := 0
	c := a / b
	fmt.Println(c)
	*/
	/* var nilMap map[int]string

	nilMap[2] = "govno"

	slice := []int{1, 2, 3}

	fmt.Println(slice[4])
	*/

	panic("balls")

}
