package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	dev := GetEnvInfo("SHOP_DEV")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if dev {
		configFileName = fmt.Sprintf("user-web/%s-dev.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("[Config] 配置加载异常:%s", err.Error())
	}

	//if err := v.Unmarshal(global.)

}
