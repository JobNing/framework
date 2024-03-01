package grpc

import (
	"fmt"
	"github.com/JobNing/framework/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
	"log"
	"net"
)

type Config struct {
	App struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	} `yaml:"app"`
}

func getConfig(serviceName string) (*Config, error) {
	configInfo, err := config.GetConfig("DEFAULT_GROUP", serviceName)
	if err != nil {
		return nil, err
	}
	cnf := new(Config)
	err = yaml.Unmarshal([]byte(configInfo), cnf)
	if err != nil {
		return nil, err
	}
	return cnf, nil
}

func RegisterGRPC(serviceName string, register func(s *grpc.Server)) error {
	cof, err := getConfig(serviceName)
	if err != nil {
		return err
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", cof.App.Ip, cof.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	//反射接口支持查询
	reflection.Register(s)
	register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return err
}
