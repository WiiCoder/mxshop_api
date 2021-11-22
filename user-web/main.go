package main

import (
	"flag"
	"fmt"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/initialize"

	"go.uber.org/zap"
)

func main() {
	ENV := flag.String("e", "dev", "ip地址")

	flag.Parse()
	fmt.Println("Env:", *ENV)

	// 1.初始化日志 logger
	initialize.InitLogger()
	// 2. 初始化配置文件
	initialize.InitConfig(*ENV)
	// 3.初始化路由 router
	Router := initialize.InitRouter()
	// 4.初始化翻译器（对gin框架响应的中文内容）
	if err := initialize.InitTranslation("zh"); err != nil {
		panic(err)
	}
	// 5.初始化服务连接
	initialize.InitSrvConn()
	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	zap.S().Debugf("[%s] 启动服务，端口：%d", global.ServerConfig.Name, global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Fatalf("[%s] 启动失败：%s", global.ServerConfig.Name, err.Error())
	}

}
