package hc

import (
	"fmt"
	"net/http"
	"syscall"

	"github.com/fighthorse/redisAdmin/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/fighthorse/redisAdmin/pkg/thirdpart/jpillora/overseer"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/v3/net"
	// "github.com/shirou/gopsutil/v3/process"
)

func Hc(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func Index(c *gin.Context) {
	name := c.Query("name")
	firstname := c.DefaultQuery("firstname", "Guest")
	log.Info(c.Request.Context(), "hc", log.Fields{"kjj": "kkk", "firstname": firstname})
	c.String(http.StatusOK, "Hello %s : %s", name, firstname)
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
		c.String(http.StatusOK, "Restart:%s:%d\n", token, pid)
		overseer.Restart()
		return
	}
	c.String(http.StatusOK, "Restart:%s", token)
}

func Stop(c *gin.Context) {
	pid := syscall.Getpid()
	c.String(http.StatusOK, "STOP %d", pid)
}
