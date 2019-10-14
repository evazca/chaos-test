package chaos_master

import (
	master "github.com/evazca/chaos-test/chaos-master"
	"github.com/evazca/chaos-test/common/log"
	"os"
)

func main() {
	log.Log()
	addr := ":4499"
	server := master.NewServer(addr,"./conf.yaml")
	err := server.Start()
	if err != nil {
		log.CommonLogger.Error(err)
		os.Exit(1)
	}
}