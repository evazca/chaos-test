package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/evazca/chaos-test/common/grpc"
	"github.com/evazca/chaos-test/common/pb"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3{
		help()
		exit(false)
	}
	addr := os.Args[1]
	rpc := os.Args[2]
	client, err:= grpc.NewMasterClient(addr)
	if err != nil {
		fmt.Println("create grpc error addr is " + addr, err)
		exit(false)
	}

	switch rpc {
	case "pr":
		done(prepare(client))
	case "rv":
		done(revoke(client))
	case "ri":
		done(randomInject(client))
	case "ts":
		done(twoSplit(client))
	case "in":
		done(isolateNode(client))
	case "no":
		args := os.Args[3:]
		done(networkOperate(client, args))
	case "rno":
		args := os.Args[3:]
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
	rpc:  pr => prepare, rv => revoke, ri => randomInject, ts => twoSplit, in => isolateNode, no => networkOperate, rno => revokeNetworkOperate
	pr: agent-cli pr
	rv: agent-cli rv
	no: agent-cli no [networkOperateMark in json form]
	rno: agent-cli rno [ip:mark string in json form]
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

func prepare(client pb.MasterClient) (string,bool){
	resp, err := client.Prepare(getContext(), &pb.PrepareMasterReq{})
	var result string
	ok := false
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.Result == false {
		result = "prepare failed"
	}else {
		result = "ok"
		ok = true
	}
	return  result,ok
}

func revoke(client pb.MasterClient) (string,bool){
	var result string
	ok := false
	resp, err := client.Revoke(getContext(), &pb.RevokeMasterReq{})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.Result == false {
		result = "revoke failed"
	}else {
		result = "ok"
		ok = true
	}
	return  result, ok
}

func randomInject(client pb.MasterClient) (string,bool){
	var result string
	ok := false
	resp, err := client.RandomInject(getContext(),&pb.RandomInjectReq{})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.Result == false {
		result = "random inject failed"
	}else {
		result = "ok"
		ok = true
	}
	return  result, ok
}

func twoSplit(client pb.MasterClient) (string,bool){
	var result string
	ok := false
	resp, err := client.TwoSplit(getContext(),&pb.TwoSplitReq{})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.Result == false {
		result = "two split failed"
	}else {
		markBytes,jerr := json.Marshal(resp.NetworkOperateMark)
		if jerr != nil{
			result = "can not cast mark to jsonStr seems impossilbe " + jerr.Error()
		}else {
			result = "ok marks are " + string(markBytes)
			ok = true
		}
	}
	return  result, ok
}

func isolateNode(client pb.MasterClient) (string,bool){
	var result string
	ok := false
	resp, err := client.IsolateNode(getContext(),&pb.IsolateNodeReq{})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.Result == false {
		result = "isolate node failed"
	}else {
		markBytes,jerr := json.Marshal(resp.NetworkOperateMark)
		if jerr != nil{
			result = "can not cast mark to jsonStr seems impossilbe " + jerr.Error()
		}else {
			result = "ok marks are " + string(markBytes)
			ok = true
		}
	}
	return  result, ok
}

func networkOperate(client pb.MasterClient, args []string) (string,bool) {
	var result string
	var ok bool
	if len(args) < 1 {
		return "args less than expect ", false
	}
	var marks []*pb.NetworkOperateMark
	jerr := json.Unmarshal([]byte(args[0]), &marks)
	if jerr != nil {
		return "networkOperateMark can not cast to struct " + jerr.Error(), false
	}
	resp, err := client.NetworkOperate(getContext(), &pb.NetworkOperateMasterReq{NetworkOperateMark:marks})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.Result == false {
		result = "networkOperate failed"
	}else {
		markBytes,jerr := json.Marshal(resp.NetworkOperateMark)
		if jerr != nil{
			result = "can not cast mark to jsonStr seems impossilbe " + jerr.Error()
		}else {
			result = "ok marks are " + string(markBytes)
			ok = true
		}
	}
	return  result, ok
}

func revokeNetworkOperate(client pb.MasterClient, args []string) (string,bool) {
	var result string
	var ok bool
	if len(args) < 1 {
		return "args less than expect", false
	}
	var ipMarks []string
	jerr := json.Unmarshal([]byte(args[0]), &ipMarks)
	if jerr != nil {
		return "marks can not cast to []string " + jerr.Error(), false
	}
	resp, err := client.RevokeNetworkOperate(getContext(), &pb.NetworkRevokeOperateMasterReq{IpMarks:ipMarks})
	if err != nil {
		result = "grpc error " + err.Error()
	}else if resp.Result == false {
		result = "revokeNetworkOperate failed"
	}else {
		result = "ok"
		ok = true
	}
	return  result, ok
}