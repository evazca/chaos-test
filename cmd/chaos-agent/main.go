package chaos_agent

import (
	agent "github.com/evazca/chaos-agent/chaos-agent/server"
	"os"
)

func main() {
	addr := ":4399"
	server := agent.NewServer(addr)
	err := server.Start()
	if err != nil {
		os.Exit(1)
	}
}