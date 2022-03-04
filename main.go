package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/fighthorse/redisAdmin/controller"
	"github.com/fighthorse/redisAdmin/pkg/conf"
	"github.com/fighthorse/redisAdmin/pkg/log"
	"github.com/fighthorse/redisAdmin/pkg/middleware"
	"github.com/fighthorse/redisAdmin/pkg/redis"
	"github.com/fighthorse/redisAdmin/pkg/thirdpart/jpillora/overseer"
	"github.com/fighthorse/redisAdmin/pkg/thirdpart/trace_redis"
	"github.com/fighthorse/redisAdmin/pkg/trace"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/spf13/viper"
)

// go build -gcflags '-m -m -m -m -l' main.go
// go tool compile -S main.go

var (
	g errgroup.Group

	v = viper.New()

	env = flag.String("env", "local", "config file name")
	//graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")

	graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")

	configpath = "./config"

	overss overseer.State

	// BuildID is compile-time variable
	BuildID = "0"
)

func readConfig(fileName, filePath string) {

	v.SetConfigName(fileName)
	v.AddConfigPath(filePath)

	// 找到并读取配置文件并且 处理错误读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := conf.Init(v); err != nil {
		panic(err)
	}

	// 初始化数据
	// redis-cfg
	trace_redis.InitCfg(conf.GConfig.Redis)

	//trace
	trace.Init()

	// log
	log.Init()

	// redis
	redis.Init()
}

func main() {
	flag.Parse()

	// 如果env使用的是绝对路径，则configpath为路径，env为文件名
	if filepath.IsAbs(*env) {
		configpath, *env = filepath.Split(*env)
	}

	readConfig(*env, configpath)
	fmt.Println("RUN Env:" + *env)
	//if *env == "prd" {
	gin.SetMode(gin.ReleaseMode)
	//}else{
	//    gin.SetMode(gin.DebugMode)
	//}

	// Creates a router without any middleware by default
	r := gin.New()
	// Global middleware
	// Recovery middleware
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	//instrument api count
	r.Use(middleware.Instrument) // defer
	// middle
	r.Use(middleware.Trace)
	//access log
	r.Use(middleware.AccessLogging) //defer

	// 初始化api依赖的各个模块
	if err := controller.Init(r); err != nil {
		panic(err)
	}
	StartListenServer(r)
}

func StartListenServer(r *gin.Engine) {
	srv := conf.GConfig.Transport
	fmt.Println("RUN Listen port " + srv.HTTP.Addr)
	//fmt.Println("Now Inner Server Listen:"+ srv.InnerHTTP.Addr)
	// 如果开启了 必须指定地址
	//var addresses []string
	//addresses = append(addresses, srv.HTTP.Addr)
	////addresses = append(addresses, srv.InnerHTTP.Addr)
	//sort.Strings(addresses)

	prog := func(state overseer.State) {
		fmt.Printf("app#%s (%s) listening... \n", BuildID, state.ID)
		err := http.Serve(state.Listener, r)
		fmt.Printf("app#%s (%s) exiting...\n", BuildID, state.ID)
		if err != nil {
			fmt.Printf("error exiting...%v \n", err)
		}
	}
	// kill -HUP pid
	overseer.Run(overseer.Config{
		Required:            true,
		Program:             prog,
		Address:             srv.HTTP.Addr,
		Addresses:           nil, // 二选一 port
		TerminateTimeout:    10 * time.Second,
		MinFetchInterval:    0,
		PreUpgrade:          nil,
		Debug:               false, //display log of overseer actions
		NoWarn:              false,
		NoRestart:           false,
		NoRestartAfterFetch: false,
		Fetcher:             nil,
	})

}
