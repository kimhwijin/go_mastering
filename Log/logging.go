package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	// 새로운 로깅 객체를 프로그램 파일이름으로 지정해서 만든다.
	//로그의 종류(local7)와 수준(info)을 지정해서 작성되는 로그는 이 수준을 따른다.
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	// 로그를 로그 서버에 보냄.
	log.Println("LOG_INFO + Log_LOCAL7: Logging in Go!")
	fmt.Println("Will you see this?")

}
