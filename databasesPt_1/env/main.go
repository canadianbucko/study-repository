package main

import (
	"os"
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	outputCount := os.Getenv("OUTPUT_COUNT")
	if outputCount == "" {
		pp.Println("Переменная окружения OUTPUT_COUNT должна быть задана")
	}
	convertedOutputCount, err := strconv.Atoi(outputCount)
	if err != nil {
		pp.Println("Ошибка преобразования OUTPUT_COUNT в целое число")
	}
	iteration := 1
	for convertedOutputCount > 0 {
		pp.Println("iteration #", iteration)
		iteration++
		convertedOutputCount--
	}

}
