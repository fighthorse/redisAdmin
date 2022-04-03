package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"time"

	"github.com/fighthorse/redisAdmin/component"
	"github.com/fighthorse/redisAdmin/component/conf"
	"github.com/fighthorse/redisAdmin/component/middleware"
	"github.com/fighthorse/redisAdmin/component/thirdpart/jpillora/overseer"
	"github.com/fighthorse/redisAdmin/controller"
	"github.com/fighthorse/redisAdmin/internal/pkg/httpserver"
	"github.com/fighthorse/redisAdmin/internal/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

// go build -gcflags '-m -m -m -m -l' main.go
// go tool compile -S main.go

var (
	v   = viper.New()
	env = flag.String("env", "local", "config file name")
	//graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")

	confidant = "./config"

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
}

func main() {
	flag.Parse()
	// 如果env使用的是绝对路径，则configpath为路径，env为文件名
	if filepath.IsAbs(*env) {
		confidant, *env = filepath.Split(*env)
	}
	//config init
	readConfig(*env, confidant)
	// 初始化数据
	component.InitComponent()
	// redis
	redis.Init()
	// http
	httpserver.Init()
	// start server
	StartListenServer()
}

func StartListenServer() {
	srv := conf.GConfig.Transport
	fmt.Println("RUN Listen port " + srv.HTTP.Addr)
	fmt.Println("RUN Listen inner port " + srv.InnerHTTP.Addr)

	// kill -HUP pid
	overseer.Run(overseer.Config{
		Required:            true,
		Program:             Program,
		Address:             "",
		Addresses:           []string{srv.HTTP.Addr, srv.InnerHTTP.Addr}, // 二选一 port
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

func Program(state overseer.State) {
	fmt.Printf("app#%s (%s) listening... \n", BuildID, state.ID)
	// inner
	if len(state.Listeners) >= 1 {
		InnerListener(state.Listeners[1])
	}

	err := http.Serve(state.Listeners[0], MainContainer())
	fmt.Printf("app#%s (%s) exiting...\n", BuildID, state.ID)
	if err != nil {
		fmt.Printf("error exiting...%v \n", err)
	}
}

func InnerListener(listener net.Listener) {
	go func() {
		//新的prometheus监控接口
		mux := http.NewServeMux()
		mux.HandleFunc("/metrics", func(writer http.ResponseWriter, request *http.Request) {
			promhttp.Handler().ServeHTTP(writer, request)
		})
		err := http.Serve(listener, mux)
		if err != nil {
			panic(err)
		}
	}()
}

func MainContainer() http.Handler {
	gin.SetMode(gin.ReleaseMode)
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
	return r
}

// init
//func main(){
//	// Creates a router without any middleware by default
//	r := gin.New()
//	// 初始化api
//	if err := controller.Init(r); err != nil {
//		panic(err)
//	}
//	_ = r.Run()
//}
