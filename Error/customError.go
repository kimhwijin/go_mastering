package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Input one or more floats")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("An Error")
	k := 1
	var n float64

	//입력들을 Float 로 Parsing 하고, Parsing 불가능하면 넘어감, 하나라도 있으면 진행하고, 하나도 없으면 return 한다.
	for err != nil {
		if k >= len(arguments) {
			fmt.Println("there is no floats")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
