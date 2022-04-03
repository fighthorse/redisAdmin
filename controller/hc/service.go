package hc

import (
	"fmt"
	"net/http"
	"os/exec"
	"syscall"
	"time"

	"github.com/fighthorse/redisAdmin/component/conf"
	"github.com/fighthorse/redisAdmin/component/thirdpart/jpillora/overseer"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/v3/net"
	// "github.com/shirou/gopsutil/v3/process"
)

func Hc(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func Index(c *gin.Context) {
	cfg := conf.GConfig.HttpServer.SelfServiceName
	c.String(http.StatusOK, "Hello %s\n", cfg)
}

func Metrics(c *gin.Context) {
	ll := promhttp.Handler()
	ll.ServeHTTP(c.Writer, c.Request)
	c.Next()
}

func Service(c *gin.Context) {

	hostInfo, _ := host.Info()
	v, _ := mem.VirtualMemory()

	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": map[string]interface{}{
		"host": hostInfo, "mem": v,
	}})
}

func Reload(c *gin.Context) {
	token := c.Query("token")
	if token == "RedisGo" {
		pid := syscall.Getpid()
		fmt.Println("reload service Before: ", pid)
		go func() {
			time.Sleep(2 * time.Second)
			overseer.Restart()
		}()
		c.String(http.StatusOK, "Restart:%s:%d\n", token, pid)
		return
	}
	c.String(http.StatusOK, "Restart:%s", token)
}

func Stop(c *gin.Context) {
	pid := syscall.Getpid()
	go func(pid int) {
		cmd := exec.Command("kill", "-HUP", fmt.Sprintf("%d", pid))
		// send close chain
		time.Sleep(3 * time.Second)
		_ = cmd.Run()
	}(pid)
	c.String(http.StatusOK, "kill -HUP %d  \n", pid)
}

func Pid(c *gin.Context) {
	pid := syscall.Getpid()
	c.String(http.StatusOK, "PID: %d  \n", pid)
}
