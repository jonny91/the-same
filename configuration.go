package main

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Folder string `yaml:"folder"`
	Mode   int    `yaml:"mode"`
}

var (
	AppConfig Config
)

func init() {
	pflag.Int("mode", 0, "设置对比模式 0:md5 1:sh1")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
}

func ReadConfig() {
	var err error
	v := viper.New()
	err = v.BindPFlags(pflag.CommandLine)
	if err != nil {
		return
	}
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")

	//读取配置文件内容
	if err = v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err = v.Unmarshal(&AppConfig); err != nil {
		panic(err)
	}
}
