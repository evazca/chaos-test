package config

import (
	"errors"
	"github.com/evazca/chaos-test/common/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

type TaskConfig struct {
	TestCmd       string `yaml:"test-cmd"`
	ServerInstances []*ServerInstance `yaml:"server-instances"`

}

type ServerInstance struct {
	IP           string   `yaml:"ip"`
	Ports        string    `yaml:"ports"`
	InstanceName string   `yaml:"instance_name"`
	StartCmd     string   `yaml:"start_cmd"`
	DataPath     string   `yaml:"data_path"`
	//the node we are not able to reach for some reason
	IgnoreNode   bool     `yaml:"ignore"`
}

func (c *TaskConfig) DecodeFile(configPath string) error {
	bs, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.CommonLogger().Error("read file error "+ configPath, err )
		return err
	}

	err = yaml.UnmarshalStrict(bs, c)
	if err != nil {
		log.CommonLogger().Error("unmrshal file error "+ configPath, err )
		return err
	}
	return nil
}

func (s *ServerInstance) GetPorts() ([]int32, error){
	if  s.Ports == "" {
		return nil, errors.New("ports not defined")
	}
	portStrs := strings.Split(s.Ports, ",")
	ports := make([]int32, 0, len(portStrs))
	for _, portStr := range portStrs {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			log.CommonLogger().Error("can not cast to int " + portStr, err)
			return nil, err
		}else {
			ports = append(ports, int32(port))
		}
	}
	return ports, nil
}