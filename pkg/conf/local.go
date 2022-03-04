package conf

type LocalConfig struct {
	Env         string `mapstructure:"env"`
	ServiceName string `mapstructure:"service_name"`
}
