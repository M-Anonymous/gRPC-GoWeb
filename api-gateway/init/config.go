package init

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"study.com/api-gateway/internal/serviceDiscovery"
)

func init() {
	//加载配置
	loadConfig()
	//获取etcd注册中心地址
	etcdAddr := viper.GetString("etcd.addr")
	//连接etcd注册中心
	serviceDiscovery.New([]string{etcdAddr})
}

func loadConfig() error {
	log.Printf("开始加载服务配置文件... \n")
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("服务配置文件加载失败,原因是: %v! \n", err)
		panic(err)
	}
	log.Printf("服务配置文件加载完成! \n")
	return nil
}
