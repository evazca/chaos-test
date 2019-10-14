package config

import (
	"github.com/evazca/chaos-test/common/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
}

func (c *TaskConfig) DecodeFile(configPath string) error {
	bs, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.CommonLogger.Error("read file error "+ configPath, err )
		return err
	}

	err = yaml.UnmarshalStrict(bs, c)
	if err != nil {
		log.CommonLogger.Error("unmrshal file error "+ configPath, err )
		return err
	}
	return nil
}