package chaos_master

import (
	"context"
	"github.com/evazca/chaos-test/common/pb"
	"google.golang.org/grpc"
)

type MockAgent struct {

}

func (m *MockAgent) Prepare(ctx context.Context, in *pb.PrepareReq, opts ...grpc.CallOption) (*pb.PrepareResp, error){
	return &pb.PrepareResp{CommonResp: &pb.CommonResp{Result:true, Id:in.Id}}, nil
}

func (m *MockAgent) Revoke(ctx context.Context, in *pb.RevokeReq, opts ...grpc.CallOption) (*pb.RevokeResp, error){
	return &pb.RevokeResp{CommonResp: &pb.CommonResp{Result:true, Id:in.Id}}, nil
}

func (m *MockAgent) NetworkOperate(ctx context.Context, in *pb.NetworkOperateReq, opts ...grpc.CallOption) (*pb.NetworkOperateResp, error){
	return &pb.NetworkOperateResp{CommonResp: &pb.CommonResp{Result:true, Id:in.Id}, Mark: -1}, nil
}

func (m *MockAgent) RevokeNetworkOperate(ctx context.Context, in *pb.RevokeNetworkOperateReq, opts ...grpc.CallOption) (*pb.RevokeNetworkOperateResp, error){
	return &pb.RevokeNetworkOperateResp{CommonResp: &pb.CommonResp{Result:true, Id:in.Id}}, nil
}