package init

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"study.com/study-user/internal/registerCenter"
)

func init() {
	//加载配置
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	//连接etcd
	endpoints := []string{viper.GetString("etcd.addr")}
	etcdRegister, err := registerCenter.New(endpoints)
	if err != nil {
		panic(err)
	}
	//注册服务
	srvName := viper.GetString("server.name")
	srvAddr := viper.GetString("server.addr")
	err = etcdRegister.Register(srvName, srvAddr, 10)
	if err != nil {
		panic(err)
	}
}

func loadConfig() error {
	//加载配置
	log.Printf("开始读取配置文件... \n")
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("配置文件读取失败,失败原因是: %v! \n", err)
	}
	log.Printf("服务配置文件读取完成! \n")
	return nil
}
