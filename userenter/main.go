package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

/*type storedItems struct {
	Name     string
	Quantity string
} */

func main() {
	items := map[string]string{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("введите команду")

	if ok := scanner.Scan(); !ok {
		fmt.Println("что-то не так")
		return
	}
	text := scanner.Text()
	fields := strings.Fields(text) //  0 - добавить удалить 1 - чеснок или еще что 2 - количество (15)

	if len(text) == 0 {
		fmt.Println("пустой массив и падаем ретурн")
		return
	}

	cmd := fields[0]
	if cmd == "добавить" {
		fmt.Println("вы хотите добавить что-то?")
		if fields[1] != "" || fields[2] != "" {
			items[fields[1]] = fields[2]
		}
	}

	fmt.Println(items)
	pp.Println(items)

}
