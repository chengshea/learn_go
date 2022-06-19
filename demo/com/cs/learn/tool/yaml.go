package tool

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var info *viper.Viper

func initReadYaml() *viper.Viper {
	pwd, _ := os.Getwd()
	conf := viper.New()
	if strings.Index(pwd, "test") > -1 {
		pwd = strings.TrimRight(pwd, "/test")
	}
	log.Println("pwd:", pwd)
	conf.AddConfigPath(pwd)
	conf.SetConfigName("info")
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		log.Fatalf("err:%s\n", err)
	}
	log.Println("conf:", conf)
	return conf
}

func getConfig() *viper.Viper {
	if info == nil {
		log.Println("init get yaml")
		info = initReadYaml()
		log.Println("init end yaml")
	}
	return info
}
