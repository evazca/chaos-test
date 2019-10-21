package main

import (
	"encoding/json"
	"fmt"
	"github.com/evazca/chaos-test/chaos-master/config"
	"github.com/evazca/chaos-test/common/pb"
)

func main()  {
	mark := pb.NetworkOperateMark{}
	separations := make([]*pb.Separation,2)
	separations[0] = &pb.Separation{Ip:"21.68.5.157"}
	separations[1] = &pb.Separation{Ip:"21.68.5.156"}
	mark.Separation = separations
	mark.NetworkOperator = &pb.NetworkOperator{Operate:pb.NetworkOperate_Loss,Probability:100}
	bytes, err :=json.Marshal(mark)
	for _,s := range separations {
		fmt.Println(s)
	}
	if err == nil{
		fmt.Println(string(bytes))
	}
	var m pb.NetworkOperateMark
	data := `
{"separation":[{"ip":"21.68.5.157"},{"ip":"21.68.5.156"}],"networkOperator":{"operate":1,"probability":100}}
`
	jerr := json.Unmarshal([]byte(data), &m)
	if jerr != nil {
		fmt.Println(jerr)
	}else {
		fmt.Println(m)
	}
	instances1 := make([]*config.ServerInstance,1)
	instances1[0] = &config.ServerInstance{IP: "21.68.5.156"}
	instances2 := make([]*config.ServerInstance,1)
	instances2[0] = &config.ServerInstance{IP: "21.68.5.157"}
	result := make([]*pb.NetworkOperateMark,0)
	operator := &pb.NetworkOperator{Operate:pb.NetworkOperate_Loss, Probability: 100}
	separations1 := make([]*pb.Separation,0)
	separations2 := make([]*pb.Separation,0)
	for _,instance := range instances1{
		separations2 = append(separations2, &pb.Separation{Ip:instance.IP})
	}
	for _,instance := range instances2{
		separations1 = append(separations1, &pb.Separation{Ip:instance.IP})
	}
	for _,instance := range instances1{
		result = append(result, &pb.NetworkOperateMark{NetworkOperator:operator,Ip: instance.IP, Separation: separations2})
	}
	for _,instance := range instances2{
		result = append(result, &pb.NetworkOperateMark{NetworkOperator:operator,Ip: instance.IP, Separation: separations1})
	}
	fmt.Println(result)
}
