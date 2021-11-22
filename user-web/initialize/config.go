package initialize

import (
	"fmt"
	"mxshop_api/user-web/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig(env string) {
	configFileName := fmt.Sprintf("config-%s.yaml", env)

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("[Config] 配置加载异常:%s", err.Error())
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Panicf("[Config] 配置解析异常:%s", err.Error())
	}

	zap.S().Info("[Config] 配置加载成功...")
	zap.S().Infof("[Config] 配置信息：%v", global.ServerConfig)

	// 动态监控配置文件
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("[Config] 配置信息变更: %v", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
	})
}
