package conf

type LocalConfig struct {
	Env         string `mapstructure:"env"`
	ServiceName string `mapstructure:"service_name"`
}

type AmapServer struct {
	Key string `mapstructure:"key"`
}
