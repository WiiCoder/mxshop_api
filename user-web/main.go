package main

import (
	"fmt"
	"mxshop_api/user-web/initialize"

	"go.uber.org/zap"
)

func main() {
	// 1.初始化日志 logger
	initialize.InitLogger()
	// 2. 初始化配置文件
	initialize.InitConfig()
	// 3.初始化路由 router
	Router := initialize.InitRouter()
	// 4.初始化翻译器（对gin框架响应的中文内容）
	if err := initialize.InitTranslation("zh"); err != nil {
		panic(err)
	}

	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	zap.S().Debugf("[User-Web] 启动服务，端口：%d", 8081)
	if err := Router.Run(fmt.Sprintf(":%d", 8081)); err != nil {
		zap.S().Fatalf("[User-Web] 启动失败：%s", err.Error())
	}

}
