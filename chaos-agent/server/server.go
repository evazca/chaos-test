package server

import (
	"context"
	"github.com/evazca/chaos-agent/common/pb"
	netMonkey "github.com/evazca/chaos-agent/monkey/network"
	"google.golang.org/grpc"
	"net"
	"sync/atomic"
)

func NewServer(addr string) *Server {
	server := Server{addr: addr, netMonkey: &netMonkey.NetMonkey{}}
	return &server
}

type Server struct {
	inProcess int32    // 0 not in process, 1 in process, 2 stopping
	id        string
	addr      string
	listener  net.Listener
	svr       *grpc.Server
	netMonkey *netMonkey.NetMonkey
}

func (s *Server) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.svr = grpc.NewServer()
	pb.RegisterAgentServer(s.svr, s)
	err = s.svr.Serve(s.listener)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Prepare(ctx context.Context, in *pb.PrepareReq) (*pb.PrepareResp, error){
	if ! atomic.CompareAndSwapInt32(&s.inProcess, 0, 1) && s.id != in.Id {
		return &pb.PrepareResp{CommonResp: &pb.CommonResp{Result:false,ErrorType: pb.ErrorType_InProcessError,Id: s.id}},nil
	}
	s.id = in.Id
	return &pb.PrepareResp{CommonResp: &pb.CommonResp{Result:true,ErrorType: pb.ErrorType_NoError,Id: s.id}},nil
}

func (s *Server) Revoke(ctx context.Context, in *pb.RevokeReq) (*pb.RevokeResp, error){
	if !in.Force && s.id != in.Id {
		return  &pb.RevokeResp{CommonResp: &pb.CommonResp{Result:false,ErrorType: pb.ErrorType_UnMatchIdError,Id: s.id}},nil
	}
	//TODO better add a latch here
	if !in.Force && !atomic.CompareAndSwapInt32(&s.inProcess, 1, 2) {
		return  &pb.RevokeResp{CommonResp: &pb.CommonResp{Result:false,ErrorType: pb.ErrorType_NotInProcessError,Id: s.id}},nil
	}
	err := s.netMonkey.Revoke()
	if err != nil {
		return  &pb.RevokeResp{CommonResp: &pb.CommonResp{Result:false,ErrorType: pb.ErrorType_UnknownError,Id: s.id}},nil
	}
	s.id = ""
	atomic.CompareAndSwapInt32(&s.inProcess, 2, 0)
	return &pb.RevokeResp{CommonResp: &pb.CommonResp{Result:true,ErrorType: pb.ErrorType_NoError,Id: s.id}},nil
}

func (s *Server) PrepareNetworkInject(ctx context.Context, in *pb.PrepareNetworkInjectReq) (*pb.PrepareNetworkInjectResp, error) {
	if in.Id != s.id {
		return  &pb.PrepareNetworkInjectResp{CommonResp: &pb.CommonResp{Result: false, ErrorType: pb.ErrorType_UnMatchIdError,Id: s.id}}, nil
	}
	err:= s.netMonkey.Prepare()
	if err != nil {
		return  &pb.PrepareNetworkInjectResp{CommonResp: &pb.CommonResp{Result: false, ErrorType: pb.ErrorType_UnknownError,Id: s.id}}, nil
	}
	return  &pb.PrepareNetworkInjectResp{CommonResp: &pb.CommonResp{Result: true, ErrorType: pb.ErrorType_NoError,Id: s.id}}, nil
}
func (s *Server) RevokeNetworkInject(ctx context.Context, in *pb.RevokeNetworkInjectReq) (*pb.RevokeNetworkInjectResp, error) {
	if in.Id != s.id {
		return  &pb.RevokeNetworkInjectResp{CommonResp: &pb.CommonResp{Result: false, ErrorType: pb.ErrorType_UnMatchIdError,Id: s.id}}, nil
	}
	err := s.netMonkey.Revoke()
	if err != nil {
		return  &pb.RevokeNetworkInjectResp{CommonResp: &pb.CommonResp{Result: false, ErrorType: pb.ErrorType_UnknownError,Id: s.id}}, nil
	}
	return  &pb.RevokeNetworkInjectResp{CommonResp: &pb.CommonResp{Result: true, ErrorType: pb.ErrorType_NoError,Id: s.id}}, nil
}
func (s *Server) NetworkOperate(ctx context.Context, in *pb.NetworkOperateReq) (*pb.NetworkOperateResp, error) {
	if in.Id != s.id {
		return  &pb.NetworkOperateResp{CommonResp: &pb.CommonResp{Result: false, ErrorType: pb.ErrorType_UnMatchIdError,Id: s.id}}, nil
	}
	mark, err := s.netMonkey.OperateNodes(in.Separation, in.Operator)
	if err != nil {
		return  &pb.NetworkOperateResp{CommonResp: &pb.CommonResp{Result: false, ErrorType: pb.ErrorType_UnknownError,Id: s.id}}, nil
	}
	return  &pb.NetworkOperateResp{CommonResp: &pb.CommonResp{Result: true, ErrorType: pb.ErrorType_NoError,Id: s.id},Mark:mark}, nil
}
func (s *Server) RevokeNetworkOperate(ctx context.Context, in *pb.RevokeNetworkOperateReq) (*pb.RevokeNetworkOperateResp, error) {
	if in.Id != s.id {
		return  &pb.RevokeNetworkOperateResp{CommonResp: &pb.CommonResp{Result: false, ErrorType: pb.ErrorType_UnMatchIdError,Id: s.id}}, nil
	}
	err := s.netMonkey.RevokeNodes(in.Mark)
	if err != nil {
		return  &pb.RevokeNetworkOperateResp{CommonResp: &pb.CommonResp{Result: false, ErrorType: pb.ErrorType_UnknownError,Id: s.id}}, nil
	}
	return  &pb.RevokeNetworkOperateResp{CommonResp: &pb.CommonResp{Result: true, ErrorType: pb.ErrorType_NoError,Id: s.id}}, nil
}
