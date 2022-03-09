package main

import (
	"os"
	"io"
)
/* 
go run stdError.go 2>/tmp/stdError
go run stdError.go 2>/dev/null
go run stdError.go >/tmp/output 2>&1
*/
func main() {
	str := ""
	arguments := os.Args
	if len(os.Args) == 1 {
		str = "no args"
	} else {
		str = arguments[1]
	}
	io.WriteString(os.Stdout, "standard out\n")
	io.WriteString(os.Stderr, str)
	io.WriteString(os.Stderr, "\n")
}