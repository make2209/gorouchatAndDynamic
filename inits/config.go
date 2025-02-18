package inits

import (
	"fmt"
	"github.com/spf13/viper"
)

type Mysql struct {
	User   string
	Passwd string
	Hort   string
	Port   string
	Data   string
}
type Redis struct {
	Addr   string
	Passwd string
	Data   int
}
type AliYun struct {
	AccessKeyID     string
	AccessKeySecret string
}
type AllConf struct {
	Mysql
	Redis
	AliYun
}

var ViperData AllConf

func InitViper() {
	viper.SetConfigFile("D:\\GoWork\\src\\zk0212\\dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&ViperData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Viper Data:%#v\n", ViperData)
}
