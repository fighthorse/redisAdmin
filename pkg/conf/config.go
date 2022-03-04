package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	GConfig = &Config{}
)

type Config struct {
	Env         string      `mapstructure:"env"`
	Commandline Commandline `mapstructure:"commandline"`
	Bootstrap   Bootstrap   `mapstructure:"bootstrap"`
	Transport   Transport   `mapstructure:"transport"`
	Log         Log         `mapstructure:"log"`
	Trace       Trace       `mapstructure:"trace"`
	Redis       []Redis     `mapstructure:"redis"`
	Server      []Server    `mapstructure:"server"`
	Nsq         Nsq         `mapstructure:"nsq"`
	Mysql       []Mysql     `mapstructure:"mysql"`
	OSS         OSS         `mapstructure:"oss"`
	LoginUser   []LoginUser `mapstructure:"login_user"`
	LocalConfig LocalConfig `mapstructure:"config"`
}

func Init(v *viper.Viper) (err error) {
	err = v.Unmarshal(GConfig)
	if err != nil {
		return err
	}
	fmt.Println("Run ServiceName:", ServiceName())
	return
}

func ServiceName() string {
	return GConfig.LocalConfig.ServiceName
}
