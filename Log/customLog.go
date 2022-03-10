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
	LOGFILE += "/GoMastering/go_mastering/Log/log/customLog.log"

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	//Lstd : Last date, Last time 로그에 현재 날짜와 시작을 같이 저장한다.
	cLog := log.New(f, "customLog ", log.LstdFlags)
	cLog.SetFlags(log.LstdFlags)
	cLog.Println("Hello Log!")
	cLog.Println("Another Log entry!")
	//customLog 2022/03/10 13:29:03 Hello Log!
	//customLog 2022/03/10 13:29:03 Another Log entry!

}
