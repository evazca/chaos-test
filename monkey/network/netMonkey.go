package network

import (
	"errors"
	"fmt"
	"github.com/evazca/chaos-test/common/log"
	"github.com/evazca/chaos-test/common/pb"
	"os/exec"
	"strconv"
	"sync"
	"sync/atomic"
)

type NetMonkey struct {
	mark int32
	separations sync.Map
	operators sync.Map
}

//TODO check the cmd result more carefully

func (n *NetMonkey) Prepare() (error) {
	//TODO it is not safe to del without check if there are some config for reason
	cmd := "sudo tc qdisc del dev eth0 root"
	log.CommonLogger.Debug(cmd)
	err := exec.Command(cmd).Run()
	if err != nil {
		log.CommonLogger.Error("exec cmd "+cmd, err)
		return  err
	}
	cmd = "sudo tc qdisc add dev eth0 handle 1: root htb"
	log.CommonLogger.Debug(cmd)
	err = exec.Command(cmd).Run()
	if err != nil {
		log.CommonLogger.Error("exec cmd "+cmd, err)
		return err
	}
	return nil
}

func (n *NetMonkey) Revoke() (error){
	cmd := "sudo tc qdisc del dev eth0 root"
	log.CommonLogger.Debug(cmd)
	err := exec.Command(cmd).Run()
	if err != nil {
		log.CommonLogger.Error("exec cmd "+cmd, err)
		return err
	}
	cmd = "sudo iptables -F"
	log.CommonLogger.Debug(cmd)
	err = exec.Command(cmd).Run()
	if err != nil {
		log.CommonLogger.Error("exec cmd "+cmd, err)
		return err
	}
	n.mark = 15
	n.separations = sync.Map{}
	n.operators = sync.Map{}
	return err
}

func (n *NetMonkey) OperateNodes(separations  []*pb.Separation, operator *pb.NetworkOperator) (int32,error) {
	mark := atomic.AddInt32(&n.mark, 1)
	cmd := fmt.Sprintf("sudo tc class add dev eth0 parent 1: classid 1:%v htb rate 1000Mbps", mark)
	log.CommonLogger.Debug(cmd)
	err := exec.Command(cmd).Run()
	if err != nil {
		log.CommonLogger.Error("exec cmd "+cmd, err)
		return mark, err
	}
	cmd = fmt.Sprintf("tc filter add dev eth0 parent 1:0 prio 1 protocol ip handle %v fw flowid 1:%v", mark, mark)
	log.CommonLogger.Debug(cmd)
	err = exec.Command(cmd).Run()
	if err != nil {
		log.CommonLogger.Error("exec cmd "+cmd, err)
		return mark, err
	}
	op := ""
	switch operator.Operate {
	case pb.NetworkOperate_Delay:
		op = "delay"
	case pb.NetworkOperate_Loss:
		op = "loss"
	default:
		return mark, errors.New("unknown operator")
	}
	cmd = fmt.Sprintf("tc qdisc add dev eth0 parent 1:%v handle %v netem %v %v%%", mark, mark, op, operator.Probability)
	log.CommonLogger.Debug(cmd)
	err = exec.Command(cmd).Run()
	if err != nil {
		log.CommonLogger.Error("exec cmd "+cmd, err)
		return mark, err
	}
	return mark, n.runIptableCmd(separations, operator, mark, false)

}

func (n *NetMonkey) runIptableCmd(separations []*pb.Separation, operator *pb.NetworkOperator , mark int32, revoke bool) error {
	n.separations.Store(mark, separations)
	n.operators.Store(mark, operator)
	for _, separation := range separations {
		op := "-A"
		if revoke {
			op = "-D"
		}
		ip := separation.Ip
		port := separation.Port
		flag := separation.Flag
		dst, dstPort, tcpFlag, protocol := "", "", "", ""
		if ip == "" {
			continue
		}
		dst = fmt.Sprintf("-d \"%v\"", ip)
		if port != 0 {
			dstPort = "--dport " + strconv.Itoa(int(port))
		}
		if flag != "" {
			protocol = "-p tcp"
			tcpFlag = "--tcp-flags " + tcpFlag
		}
		cmd := fmt.Sprintf("iptables %v OUTPUT -t mangle %v %v %v %v -j MARK --set-mark %v", op, protocol, tcpFlag, dst, dstPort, mark)
		log.CommonLogger.Debug(cmd)
		err := exec.Command(cmd).Run()
		if err != nil {
			log.CommonLogger.Error("exec cmd "+cmd, err)
			return err
		}
	}
	return nil
}

func (n *NetMonkey) RevokeNodes(mark int32) error {
	defer func(){
		n.separations.Delete(mark)
		n.operators.Delete(mark)
	}()
	separationsI, ok := n.separations.Load(mark)
	if !ok {
		return nil
	}
	separations, ok := separationsI.([]*pb.Separation)
	if !ok {
		log.CommonLogger.Error(separationsI)
		return errors.New("unknown interface cast error")
	}
	operatorI,ok := n.operators.Load(mark)

	if !ok {
		return nil
	}
	operator, ok := operatorI.(*pb.NetworkOperator)
	if !ok {
		log.CommonLogger.Error(operatorI)
		return errors.New("unknown interface cast error")
	}

	return n.runIptableCmd(separations, operator, mark, true)
}