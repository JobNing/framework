package consul

import (
	"fmt"
	"github.com/JobNing/framework/config"
	"github.com/google/uuid"
	capi "github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
	"strconv"
)

type ConsulConfig struct {
	Consul struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	} `yaml:"consul"`
}

func getConfig(nacosGroup, serviceName string) (*ConsulConfig, error) {
	cnf, err := config.GetConfig(nacosGroup, serviceName)
	if err != nil {
		return nil, err
	}

	consulCnf := new(ConsulConfig)
	err = yaml.Unmarshal([]byte(cnf), consulCnf)
	if err != nil {
		return nil, err
	}

	return consulCnf, err
}

func AgentHealthService(serviceName string) (string, error) {
	client, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		return "", err
	}
	sr, info, err := client.Agent().AgentHealthServiceByName(serviceName)
	if err != nil {
		return "", err
	}
	if sr != "passing" {
		return "", fmt.Errorf("is not have health service")
	}
	return fmt.Sprintf("%v:%v", info[0].Service.Address, info[0].Service.Port), nil
}

func ServiceRegister(nacosGroup, serviceName string, address, port string) error {
	cof, err := getConfig(nacosGroup, serviceName)
	if err != nil {
		return err
	}
	client, err := capi.NewClient(&capi.Config{
		Address: fmt.Sprintf("%v:%v", cof.Consul.Ip, cof.Consul.Port),
	})
	if err != nil {
		return err
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	return client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    "user",
		Tags:    []string{"GRPC"},
		Port:    portInt,
		Address: address,
		//Check: &capi.AgentServiceCheck{
		//	GRPC:                           fmt.Sprintf("%v:%v", address, port),
		//	Interval:                       "5s",
		//	DeregisterCriticalServiceAfter: "10s",
		//},
	})
}
