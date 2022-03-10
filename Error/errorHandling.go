package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Same integer error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err != nil {
		log.Println(err)
		os.Exit(10)
	}
	fmt.Println("normal return")

	err = returnError(10, 10)
	if err != nil {
		/*
			normal return
			panic: Same integer error in returnError() function!

			goroutine 1 [running]:
			main.main()
					/Error/errorHandling.go:29 +0xb4
			exit status 2
		*/
		panic(err)
		os.Exit(10)
	}
	fmt.Println("normal return")

}
