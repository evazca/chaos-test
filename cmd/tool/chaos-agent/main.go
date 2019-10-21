package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/evazca/chaos-test/common/grpc"
	"github.com/evazca/chaos-test/common/pb"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 4{
		help()
		exit(false)
	}
	addr := os.Args[1]
	rpc := os.Args[2]
	args := os.Args[3:]
	client, err:= grpc.NewAgentClient(addr)
	if err != nil {
		fmt.Println("create grpc error addr is " + addr, err)
		exit(false)
	}

	switch rpc {
	case "pr":
		done(prepare(client, args))
	case "rv":
		done(revoke(client, args))
	case "no":
		done(networkOperate(client, args))
	case "rno":
		done(revokeNetworkOperate(client, args))
	default:
		help()
		exit(false)
	}
	exit(true)
}

func help()  {
	var help = `
	used for debug or special job like force revoke
	sample: agent-cli addr rpc args
	rpc:  pr => prepare, rv => revoke, no => networkOperate, rno => revokeNetworkOperate
	pr: pr id
	rv: rv id force
	no: no id [networkOperateMark in json form]
	rno: rno id mark
	`
	fmt.Println(help)
}

func getContext() context.Context {
	ctx,_ := context.WithTimeout(context.Background(), time.Second * 5)
	return ctx
}

func exit(ok bool)  {
	if ok {
		os.Exit(0)
	}else {
		os.Exit(-1)
	}
}

func done(result string, ok bool) {
	fmt.Println(result)
	exit(ok)

}

func prepare(client pb.AgentClient, args []string) (string,bool){
	id := args[0]
	resp, err := client.Prepare(getContext(), &pb.PrepareReq{Id: id})
	var result string
	ok := false
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.CommonResp.Result == false {
		result = "prepare failed current id is " + resp.CommonResp.Id
	}else {
		result = "ok"
		ok = true
	}
	return  result,ok
}

func revoke(client pb.AgentClient, args []string) (string,bool){
	var result string
	ok := false
	if len(args) < 2 {
		return "args less than expect ", false
	}
	id := args[0]
	force,berr := strconv.ParseBool(args[1])
	if berr != nil{
		return "force must be bool string, but force is " + args[1], false
	}
	resp, err := client.Revoke(getContext(), &pb.RevokeReq{Id: id,Force:force})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.CommonResp.Result == false {
		result = "revoke failed current id is " + resp.CommonResp.Id
	}else {
		result = "ok"
		ok = true
	}
	return  result, ok
}

func networkOperate(client pb.AgentClient, args []string) (string,bool) {
	var result string
	var ok bool
	if len(args) < 2 {
		return "args less than expect ", false
	}
	id := args[0]
	var mark pb.NetworkOperateMark
	jerr := json.Unmarshal([]byte(args[1]), &mark)
	if jerr != nil {
		return "networkOperateMark can not cast to struct " + args[1] + jerr.Error(), false
	}
	resp, err := client.NetworkOperate(getContext(), &pb.NetworkOperateReq{Id:id,NetworkOperateMark: &mark})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.CommonResp.Result == false {
		result = "networkOperate failed current id is " + resp.CommonResp.Id
	}else {
		result = "ok mark is " + strconv.Itoa(int(resp.Mark))
		ok = true
	}
	return  result, ok
}

func revokeNetworkOperate(client pb.AgentClient, args []string) (string,bool) {
	var result string
	var ok bool
	if len(args) < 2 {
		return "args less than expect", false
	}
	id := args[0]
	mark,ierr := strconv.ParseInt(args[1], 10, 32)
	if ierr != nil {
		return "mark is not int32 string", false
	}
	resp, err := client.RevokeNetworkOperate(getContext(), &pb.RevokeNetworkOperateReq{Id:id,Mark:int32(mark)})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.CommonResp.Result == false {
		result = "networkOperate failed current id is " + resp.CommonResp.Id
	}else {
		result = "ok"
		ok = true
	}
	return  result, ok
}