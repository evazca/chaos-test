package grpc

import (
	"github.com/evazca/chaos-test/common/log"
	"github.com/evazca/chaos-test/common/pb"
	"google.golang.org/grpc"
)

func NewAgentClient(target string) (pb.AgentClient,error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.CommonLogger().Error("grpc dial error "+target,err)
		return nil,err
	}

	return pb.NewAgentClient(conn),nil
}

func NewMasterClient(target string) (pb.MasterClient,error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.CommonLogger().Error("grpc dial error "+target,err)
		return nil,err
	}

	return pb.NewMasterClient(conn),nil
}