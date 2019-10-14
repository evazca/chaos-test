package chaos_agent

import (
	agent "github.com/evazca/chaos-test/chaos-agent"
	"github.com/evazca/chaos-test/common/log"
	"os"
)

func main() {
	log.Log()
	addr := ":4399"
	server := agent.NewServer(addr)
	err := server.Start()
	if err != nil {
		log.CommonLogger.Error(err)
		os.Exit(1)
	}
}