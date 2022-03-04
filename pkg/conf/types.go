package conf

type Commandline struct {
	Graceful bool `mapstructure:"graceful"`
}

type Bootstrap struct {
	Graceful bool   `mapstructure:"graceful"`
	Pid      string `mapstructure:"pid"`
	Timeout  int    `mapstructure:"timeout"`
}

type Transport struct {
	HTTP         HTTPConfig `mapstructure:"http"`
	InnerHTTP    HTTPConfig `mapstructure:"inner_http"`
	CmdInnerHTTP HTTPConfig `mapstructure:"cmd_inner_http"`
	Grpc         GrpcConfig `mapstructure:"grpc"`
}

type HTTPConfig struct {
	Addr              string  `mapstructure:"addr"`
	MaxConns          int     `mapstructure:"max_conns"`
	ReadTimeout       float64 `mapstructure:"read_timeout"`
	ReadHeaderTimeout float64 `mapstructure:"read_header_timeout"`
	WriteTimeout      float64 `mapstructure:"write_timeout"`
	IdleTimeout       float64 `mapstructure:"idle_timeout"`
}

type GrpcConfig struct {
	Addr string `mapstructure:"addr"`
}

type Log struct {
	App struct {
		FilePath string `mapstructure:"file_path"`
		Level    string `mapstructure:"level"`
	} `mapstructure:"app"`

	Access struct {
		FilePath string `mapstructure:"file_path"`
	} `mapstructure:"access"`
}

type Trace struct {
	ServiceName string  `mapstructure:"service_name"`
	FilePath    string  `mapstructure:"file_path"`
	Sampling    float64 `mapstructure:"sampling"`
}

type Redis struct {
	Name         string  `mapstructure:"name"`
	Addr         string  `mapstructure:"addr"`
	Pwd          string  `mapstructure:"pwd"`
	Db           float64 `mapstructure:"db"`
	DialTimeout  float64 `mapstructure:"dial_timeout"`
	ReadTimeout  float64 `mapstructure:"read_timeout"`
	WriteTimeout float64 `mapstructure:"write_timeout"`
	PoolSize     float64 `mapstructure:"pool_size"`
	MinIdleConns float64 `mapstructure:"min_idle_conns"`
	MaxRetries   float64 `mapstructure:"max_retries"`
}

type Mysql struct {
	Name   string      `mapstructure:"name"`
	Master MysqlConfig `mapstructure:"master"`
	Slave  MysqlConfig `mapstructure:"slave"`
}

type MysqlConfig struct {
	Driver         string  `mapstructure:"driver"`
	DSN            string  `mapstructure:"dsn"`
	MaxOpenConns   int32   `mapstructure:"max_open_conns"`
	MaxIdleConns   int32   `mapstructure:"max_idle_conns"`
	MaxLifeTimeout float64 `mapstructure:"max_life_timeout"`
}

type Kv struct {
	Url       string `mapstructure:"url"`
	KeyPerfix string `mapstructure:"key_perfix"`
}

type Server struct {
	Name                 string `mapstructure:"name"`
	Url                  string `mapstructure:"url"`
	DiscoveryServiceName string `mapstructure:"discovery_service_name"`
	DiscoveryServicePort int    `mapstructure:"discovery_service_port"`
	DiscoveryTag         string `mapstructure:"discovery_tag"`
	DiscoveryDC          string `mapstructure:"discovery_dc"`
	ConnectSidecar       string `mapstructure:"connect_sidecar"`
	ConnectConsul        string `mapstructure:"connect_consul"`
}

type Nsq struct {
	Producer []Producer `mapstructure:"producer"`
	Consumer []Consumer `mapstructure:"consumer"`
}

type Producer struct {
	Name           string `mapstructure:"name"`
	Addr           string `mapstructure:"addr"`
	MaxConcurrency int32  `mapstructure:"max_concurrency"`
	DialTimeout    int32  `mapstructure:"dial_timeout"`
	ReadTimeout    int32  `mapstructure:"read_timeout"`
	WriteTimeout   int32  `mapstructure:"write_timeout"`
}

type Consumer struct {
	Name         string   `mapstructure:"name"`
	Addr         string   `mapstructure:"addr"`
	Lookup       []string `mapstructure:"lookup"`
	MaxInFlight  int32    `mapstructure:"max_inflight"`
	DialTimeout  int32    `mapstructure:"dial_timeout"`
	ReadTimeout  int32    `mapstructure:"read_timeout"`
	WriteTimeout int32    `mapstructure:"write_timeout"`
}

type OSS struct {
	AK       string `mapstructure:"ak"`
	AS       string `mapstructure:"as"`
	Endpoint string `mapstructure:"endpoint"`
}

type LoginUser struct {
	UserName string `mapstructure:"user_name"`
	UserPwd  string `mapstructure:"user_pwd"`
}
