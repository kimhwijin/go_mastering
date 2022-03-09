package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f) // 새로운 묵시적 변수를 정의할때 := 를 사용함.
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

}
