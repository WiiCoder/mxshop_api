package initialize

import (
	"fmt"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	ip := global.ServerConfig.UserSrvConfig.Host
	port := global.ServerConfig.UserSrvConfig.Port
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GRPC Server] 连接【用户服务】失败")
		return
	}

	UserSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = UserSrvClient
}
