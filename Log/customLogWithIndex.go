package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	LOGFILE, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	LOGFILE += "/GoMastering/go_mastering/Log/log/customLogWithIndex.log"

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	//Lstd : Last date, Last time 로그에 현재 날짜와 시작을 같이 저장한다.
	cLog := log.New(f, "customLogWithIndex ", log.LstdFlags)
	cLog.SetFlags(log.LstdFlags | log.Lshortfile)
	// 소스코드의 줄번호까지 입력한다.
	cLog.Println("Hello Log!")
	cLog.Println("Another Log entry!")
	//customLogWithIndex 2022/03/10 13:40:40 customLogWithIndex.go:29: Hello Log!
	//customLogWithIndex 2022/03/10 13:40:40 customLogWithIndex.go:30: Another Log entry!
}
