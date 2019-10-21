package main

import (
	master "github.com/evazca/chaos-test/chaos-master"
	"github.com/evazca/chaos-test/common/log"
	"os"
	"time"
)

func main() {
	log.Log()
	addr := ":4499"
	server := master.NewServer(addr,"./conf.yaml")
	err := server.Start()
	if err != nil {
		log.CommonLogger().Error(err)
		//TODO log4go has flush problem when stop the program,use uber zap log later
		time.Sleep(time.Millisecond * 500)
		os.Exit(1)
	}
	//TODO log4go has flush problem when stop the program,use uber zap log later
	time.Sleep(time.Millisecond * 500)
	os.Exit(0)
}