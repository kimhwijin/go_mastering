package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)	
	}

	//arg[0] 은 실행가능한 파일이름
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64) //parse 한 결과, error 변수 반환
	max, _ := strconv.ParseFloat(arguments[1], 64)
	
	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min :", min)
	fmt.Println("Max :", max)
}
