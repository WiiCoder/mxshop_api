package initialize

import (
	"fmt"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/proto"

	"google.golang.org/grpc"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulConfig
	userConn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=15s",
		consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvConfig.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Errorw("[GRPC Server] 连接【用户服务】失败")
		return
	}
	UserSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = UserSrvClient
}

func InitSrvConn2() {
	// 1. 直连方式
	//ip := global.ServerConfig.UserSrvConfig.Host
	//port := global.ServerConfig.UserSrvConfig.Port
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	//if err != nil {
	//	zap.S().Errorw("[GRPC Server] 连接【用户服务】失败")
	//	return
	//}
	//
	//UserSrvClient := proto.NewUserClient(userConn)
	//global.UserSrvClient = UserSrvClient

	// 2. 注册中心获取

	config := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulConfig
	config.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Panic(err)
	}

	filterStr := fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvConfig.Name)
	data, err := client.Agent().ServicesWithFilter(filterStr)
	if err != nil {
		zap.S().Panic(err)
	}

	userSrvHost := ""
	userSrvPort := 0

	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GRPC Server] 连接【用户服务】失败")
		return
	}
	UserSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = UserSrvClient

}
