package consul

import (
	"fmt"
	capi "github.com/hashicorp/consul/api"
	"strconv"
)

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

func ServiceRegister(address, port string) error {
	client, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		return err
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	return client.Agent().ServiceRegister(&capi.AgentServiceRegistration{
		ID:      "test",
		Name:    "user",
		Tags:    []string{"GRPC"},
		Port:    portInt,
		Address: address,
		Check: &capi.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%v:%v", address, port),
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "10s",
		},
	})
}
