package main

import (
	"errors"
	"fmt"
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
	if err == nil {
		fmt.Println("returnError return normal")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError return normal")
	} else {
		// 에러 객체
		fmt.Println(err)
		// 에러내용 스트링 변환
		fmt.Println(err.Error())
	}

}
