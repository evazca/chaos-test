package chaos_master

import (
	"context"
	"errors"
	"github.com/evazca/chaos-test/chaos-master/config"
	"github.com/evazca/chaos-test/common/exec"
	grpc2 "github.com/evazca/chaos-test/common/grpc"
	"github.com/evazca/chaos-test/common/log"
	"github.com/evazca/chaos-test/common/pb"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func NewServer(addr string, configPath string) *Server {
	server := Server{addr: addr, config: config.TaskConfig{}, configPath: configPath, agentMap: make(map[string]pb.AgentClient), timeout: time.Second}
	return &server
}

type Server struct {
	inProcess       int32 // 0 not in process, 1 in process, 2 stopping
	addr            string
	listener        net.Listener
	svr             *grpc.Server
	prepared        bool
	configPath      string
	config          config.TaskConfig
	agentMap        map[string]pb.AgentClient
	timeout         time.Duration
	id              string
	hostname        string
	networkOperates sync.Map
}

func (s *Server) Start() error {
	var err error
	s.hostname, err = os.Hostname()
	if err != nil {
		log.CommonLogger().Error("hostname error ", err)
		return err
	}
	s.id = s.hostname + ":" + strconv.FormatInt(time.Now().Unix(), 10)
	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.CommonLogger().Error("listen addr error "+s.addr, err)
		return err
	}
	err = s.config.DecodeFile(s.configPath)
	if err != nil {
		return err
	}
	for _, server := range s.config.ServerInstances {
		if !server.IgnoreNode {
			s.agentMap[server.IP], err = grpc2.NewAgentClient(server.IP + ":4399")
			if err != nil {
				return err
			}
		} else {
			//every grpc call to the ignore node will return success
			s.agentMap[server.IP] = &MockAgent{}
		}
	}

	s.svr = grpc.NewServer()
	pb.RegisterMasterServer(s.svr, s)
	s.inProcess = 1
	go func() {
		time.Sleep(time.Second)
		cmdStr := s.config.TestCmd
		cmd, stdout, stderr := exec.Command(cmdStr)
		err := cmd.Run()
		log.CommonLogger().Info(stdout.String())
		if stderr.String() != "" {
			log.CommonLogger().Error(stderr.String())
		}

		if err != nil {
			log.CommonLogger().Error("test cmd exec error "+cmdStr, err)
		}
		ctx, _ := context.WithTimeout(context.Background(), s.timeout)
		err = s.revoke(ctx)
		if err != nil {
			log.CommonLogger().Error("revoke error ", err)
		}
		s.inProcess = 2
		s.Close()
	}()
	err = s.svr.Serve(s.listener)
	if err != nil {
		if s.inProcess == 2 {
			return nil
		}
		log.CommonLogger().Error("grpc serve error ", err)
		return err
	}
	return nil
}

func (s *Server) Prepare(ctx context.Context, in *pb.PrepareMasterReq) (*pb.PrepareMasterResp, error) {
	preparedIp := make([]string, 0)
	for ip, agent := range s.agentMap {
		ctx1, _ := context.WithTimeout(ctx, s.timeout)
		resp, err := agent.Prepare(ctx1, &pb.PrepareReq{Id: s.id})
		if err != nil || resp.CommonResp.Result != true {
			log.CommonLogger().Error("prepare error ip is "+ip, err, resp)
			for _, ipr := range preparedIp {
				ctx2, _ := context.WithTimeout(ctx, s.timeout)
				rspR, errR := s.agentMap[ipr].Revoke(ctx2, &pb.RevokeReq{Id: s.id})
				if errR != nil || rspR.CommonResp.Result != true {
					log.CommonLogger().Error("revoke error ip is "+ipr, errR, rspR)
				}
			}
			return &pb.PrepareMasterResp{Result: false}, nil
		}
		preparedIp = append(preparedIp, ip)
	}
	s.prepared = true
	return &pb.PrepareMasterResp{Result: true}, nil
}

func (s *Server) Revoke(ctx context.Context, in *pb.RevokeMasterReq) (*pb.RevokeMasterResp, error) {
	err := s.revoke(ctx)
	if err != nil {
		return &pb.RevokeMasterResp{Result: false}, nil
	}
	return &pb.RevokeMasterResp{Result: true}, nil
}

func (s *Server) RandomInject(ctx context.Context, in *pb.RandomInjectReq) (*pb.RandomInjectResp, error) {
	return &pb.RandomInjectResp{Result: true}, nil
}

func (s *Server) TwoSplit(ctx context.Context, in *pb.TwoSplitReq) (*pb.TwoSplitResp, error) {
	instances1, instances2, err := s.splitRegion()
	if err != nil {
		log.CommonLogger().Error("split region err", err)
		return &pb.TwoSplitResp{Result: false}, nil
	}
	_, err = s.operateSplits(ctx, s.createOperateMark(instances1, instances2))
	if err != nil {
		return &pb.TwoSplitResp{Result: false}, nil
	}
	return &pb.TwoSplitResp{Result: true, NetworkOperateMark: s.getAllOperates()}, nil
}

func (s *Server) IsolateNode(ctx context.Context, in *pb.IsolateNodeReq) (*pb.IsolateNodeResp, error) {
	instances1, instances2, err := s.splitOneNode()
	if err != nil {
		log.CommonLogger().Error("split one node err", err)
		return &pb.IsolateNodeResp{Result: false}, nil
	}
	_, err = s.operateSplits(ctx, s.createOperateMark(instances1, instances2))
	if err != nil {
		return &pb.IsolateNodeResp{Result: false}, nil
	}
	return &pb.IsolateNodeResp{Result: true, NetworkOperateMark: s.getAllOperates()}, nil
}

func (s *Server) NetworkOperate(ctx context.Context, in *pb.NetworkOperateMasterReq) (*pb.NetworkOperateMasterResp, error) {
	marks, err := s.operateSplits(ctx, in.NetworkOperateMark)
	if err != nil {
		return &pb.NetworkOperateMasterResp{Result: false}, nil
	}
	return &pb.NetworkOperateMasterResp{Result: true, NetworkOperateMark: marks}, nil
}

func (s *Server) RevokeNetworkOperate(ctx context.Context, in *pb.NetworkRevokeOperateMasterReq) (*pb.NetworkRevokeOperateMasterResp, error) {
	err := s.revokeSplits(ctx, in.IpMarks)
	if err != nil {
		return &pb.NetworkRevokeOperateMasterResp{Result: false}, nil
	}
	return &pb.NetworkRevokeOperateMasterResp{Result: true}, nil
}

func (s *Server) Close() {
	err := s.listener.Close()
	if err != nil {
		log.CommonLogger().Error(err)
	}
	if s.svr != nil {
		s.svr.Stop()
	}
}

func (s *Server) revoke(ctx context.Context) error {
	if !s.prepared {
		return nil
	}
	var fail bool
	for ip, agent := range s.agentMap {
		ctx1, _ := context.WithTimeout(ctx, s.timeout)
		resp, err := agent.Revoke(ctx1, &pb.RevokeReq{Id: s.id})
		if err != nil || resp.CommonResp.Result != true {
			log.CommonLogger().Error("revoke error ip is "+ip, err, resp)
			fail = true
		}
	}
	if fail {
		return errors.New("failed to revoke for some reason")
	}
	return nil
}

func (s *Server) splitRegion() ([]*config.ServerInstance, []*config.ServerInstance, error) {
	if len(s.config.ServerInstances) < 2 {
		return nil, nil, errors.New("brain split can't happen when there is less than 2 node")
	}
	instances1 := make([]*config.ServerInstance, 0)
	instances2 := make([]*config.ServerInstance, 0)
	rand.Seed(time.Now().UnixNano())
	for _, instance := range s.config.ServerInstances {
		if rand.Intn(2) == 0 {
			instances1 = append(instances1, instance)
		} else {
			instances2 = append(instances2, instance)
		}
	}
	if len(instances1) == 0 || len(instances2) == 0 {
		//use split one node instead of rand again
		return s.splitOneNode()
	}
	return instances1, instances2, nil
}

func (s *Server) splitOneNode() ([]*config.ServerInstance, []*config.ServerInstance, error) {
	instanceCnt := len(s.config.ServerInstances)
	instances := s.config.ServerInstances
	if instanceCnt < 2 {
		return nil, nil, errors.New("brain split can't happen when there is less than 2 node")
	}
	instances1 := make([]*config.ServerInstance, 0)
	instances2 := make([]*config.ServerInstance, 0)
	rand.Seed(time.Now().UnixNano())
	randNodeNum := rand.Intn(instanceCnt)
	for n := 0; n < instanceCnt; n++ {
		if n != randNodeNum {
			instances1 = append(instances1, instances[n])
		} else {
			instances2 = append(instances2, instances[n])
		}
	}
	return instances1, instances2, nil
}

func (s *Server) createOperateMark(instances1 []*config.ServerInstance, instances2 []*config.ServerInstance) []*pb.NetworkOperateMark {
	result := make([]*pb.NetworkOperateMark, 0)
	operator := &pb.NetworkOperator{Operate: pb.NetworkOperate_Loss, Probability: 100}
	separations1 := getSeparations(instances1)
	separations2 := getSeparations(instances2)

	for _, instance := range instances1 {
		result = append(result, &pb.NetworkOperateMark{NetworkOperator: operator, Ip: instance.IP, Separation: separations1})
	}
	for _, instance := range instances2 {
		result = append(result, &pb.NetworkOperateMark{NetworkOperator: operator, Ip: instance.IP, Separation: separations2})
	}
	log.CommonLogger().Debug(*instances1[0], *instances2[0], separations1, separations2, result)
	return result
}

func (s *Server) operateSplits(ctx context.Context, splits []*pb.NetworkOperateMark) ([]*pb.NetworkOperateMark, error) {
	marks := make([]*pb.NetworkOperateMark, 0)
	for _, operateMark := range splits {
		ip := operateMark.Ip
		ctx1, _ := context.WithTimeout(ctx, s.timeout)
		resp, err := s.agentMap[ip].NetworkOperate(ctx1, &pb.NetworkOperateReq{Id: s.id, NetworkOperateMark: operateMark})
		if err != nil || resp.CommonResp.Result != true {
			log.CommonLogger().Error("operate split failed ", operateMark, err, resp)
			return nil, errors.New("operate split failed")
		}
		mark := resp.Mark
		//the mark mock client return is -1
		if mark > 0 {
			netWorkOperateMark := &pb.NetworkOperateMark{Mark: mark, Separation: operateMark.Separation, NetworkOperator: operateMark.NetworkOperator, Ip: ip}
			//mark 不够 ip:mark
			s.networkOperates.Store(ip+":"+strconv.Itoa(int(mark)), netWorkOperateMark)
			marks = append(marks, netWorkOperateMark)
		}
	}
	return marks, nil
}

func (s *Server) getAllOperates() []*pb.NetworkOperateMark {
	result := make([]*pb.NetworkOperateMark, 0)
	s.networkOperates.Range(func(key, value interface{}) bool {
		operate, ok := value.(*pb.NetworkOperateMark)
		if ok {
			result = append(result, operate)
		} else {
			log.CommonLogger().Error("cast to *pb.NetworkOperateMark error ", value)
		}
		return true
	})
	return result
}

func (s *Server) revokeSplits(ctx context.Context, ipMarks []string) error {
	for _, ipMark := range ipMarks {
		if networkMarkI, ok := s.networkOperates.Load(ipMark); ok {
			networkMark, okI := networkMarkI.(*pb.NetworkOperateMark)
			if !okI {
				log.CommonLogger().Error("cast to *pb.NetworkOperateMark error ", networkMarkI)
				return errors.New("unexpected exception")
			}
			ip := networkMark.Ip
			ctx1, _ := context.WithTimeout(ctx, s.timeout)
			resp, err := s.agentMap[ip].RevokeNetworkOperate(ctx1, &pb.RevokeNetworkOperateReq{Id: s.id, Mark: networkMark.Mark})
			if err != nil || resp.CommonResp.Result != true {
				log.CommonLogger().Error("revoke split failed ip is "+ip+"mark is "+strconv.Itoa(int(networkMark.Mark)), err, resp)
				return errors.New("revoke split failed")
			}
			s.networkOperates.Delete(ipMark)
		} else {
			return errors.New("ip:mark " + ipMark + " not exist")
		}
	}
	return nil
}

func getSeparations(instances []*config.ServerInstance) ([]*pb.Separation) {
	separations := make([]*pb.Separation, 0)
	for _, instance := range instances {
		ports, err := instance.GetPorts()
		if err != nil {
			separations = append(separations, &pb.Separation{Ip: instance.IP})
		} else {
			for _, port := range ports {
				separations = append(separations, &pb.Separation{Ip: instance.IP,Port:port})
			}
		}
	}
	return separations
}
