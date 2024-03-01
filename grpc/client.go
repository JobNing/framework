package grpc

import (
	"fmt"
	"github.com/JobNing/framework/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
)

func Client(toService string) (*grpc.ClientConn, error) {
	cnfStr, err := config.GetConfig("DEFAULT_GROUP", toService)
	if err != nil {
		return nil, err
	}
	cnf := new(Config)
	err = yaml.Unmarshal([]byte(cnfStr), &cnf)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(fmt.Sprintf("%v:%v", cnf.App.Ip, cnf.App.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
