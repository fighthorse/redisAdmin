package redis

import (
	"errors"

	"github.com/fighthorse/redisAdmin/component/self_errors"
	"github.com/fighthorse/redisAdmin/internal/service/work"
	"github.com/fighthorse/redisAdmin/protos"
	"github.com/gin-gonic/gin"
)

// SearchInit
func SearchInit(c *gin.Context) {
	data, err := work.ListRedisCfg(c)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
	return
}
func Search(c *gin.Context) {
	var db protos.SearchReq
	if err := c.ShouldBind(&db); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	data, err := work.Search(c, db)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
	return
}

func SearchKey(c *gin.Context) {
	//AddCfgReq
	var search protos.SearchKeyReq
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}

	data, err := work.HandleKey(c, search)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
	return
}

func SearchNowKey(c *gin.Context) {
	//AddCfgReq
	var search protos.SearchKeyReq
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}

	data, err := work.HandleNowKey(c, search)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
	return
}

func GetKey(c *gin.Context) {
	//AddCfgReq
	var search protos.SearchKeyReq
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}

	data, err := work.GetKey(c, search)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
	return
}

func Info(c *gin.Context) {
	data, err := work.ListRedisCfg(c)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
	return
}

func AddCfg(c *gin.Context) {
	//AddCfgReq
	var db protos.AddCfgReq
	if err := c.ShouldBind(&db); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	if db.Name == "" || db.Addr == "" {
		err := errors.New("用户名称/密码不能为空")
		c.JSON(200, self_errors.JsonErrExport(self_errors.ParamsErr, err, ""))
		return
	}
	data, err := work.AddRedisCfg(c, db)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
	return
}

// 根据类型处理
func Handle(c *gin.Context) {
	//AddCfgReq
	var search protos.SearchKeyReq
	if err := c.ShouldBind(&search); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}

	data, err := work.HandleKey(c, search)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
	return
}
