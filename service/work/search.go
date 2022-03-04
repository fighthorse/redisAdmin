package work

import (
	"errors"
	"strconv"
	"strings"

	"github.com/fighthorse/redisAdmin/pkg/redis"
	"github.com/fighthorse/redisAdmin/pkg/thirdpart/trace_redis"
	"github.com/fighthorse/redisAdmin/protos"
	"github.com/gin-gonic/gin"
)

func HandleNowKey(c *gin.Context, req protos.SearchKeyReq) (interface{}, error) {
	// get client
	db, _ := strconv.Atoi(req.Db)
	client := redis.LoadOthersDB(req.Client, db)
	if client == nil {
		return nil, errors.New("redis client create error")
	}

	if !strings.Contains(req.Key, "*") {
		return GetNoneKey(c, req, req.Key, client.Client)
	}
	// 先scan 是否只有一个
	key := req.Key
	if !strings.Contains(req.Key, "*") {
		key += "*"
	}
	keysPerFix, count := ScanRedis(client.Client, key, req.Level+1)
	if count == 1 && !strings.Contains(keysPerFix[0], "*") {
		return GetNoneKey(c, req, keysPerFix[0], client.Client)
	}

	out := make(map[string]int, len(keysPerFix))
	for _, v := range keysPerFix {
		out[v] += 1
	}

	res := make(map[string]interface{})
	res["type"] = ""
	res["keys"] = out
	res["count"] = count
	res["child"] = false
	if count == 1 && strings.Contains(keysPerFix[0], "*") {
		res["child"] = true
	}
	if count > 1 {
		res["child"] = true
	}
	res["level"] = req.Level + 1
	return res, nil
}

func GetKey(c *gin.Context, req protos.SearchKeyReq) (interface{}, error) {
	// get client
	db, _ := strconv.Atoi(req.Db)
	client := redis.LoadOthersDB(req.Client, db)
	if client == nil {
		return nil, errors.New("redis client create error")
	}
	return GetNoneKey(c, req, req.Key, client.Client)
}

func GetNoneKey(c *gin.Context, req protos.SearchKeyReq, keys string, client *trace_redis.RedisClient) (interface{}, error) {
	// handleKey
	typeInfo, err := client.Type(keys).Result()
	if err != nil {
		return nil, err
	}
	if typeInfo == "none" {
		return nil, errors.New("当前key不存在")
	}
	res := GetKeyByType(c, req, typeInfo, keys, client)
	return res, nil
}

func Search(c *gin.Context, data protos.SearchReq) (interface{}, error) {
	// get client
	db, _ := strconv.Atoi(data.Db)
	client := redis.LoadOthersDB(data.Client, db)
	if client == nil {
		return nil, errors.New("redis client create error")
	}
	match := data.Key
	if match == "" {
		match = "*"
	}
	level := data.Level
	if level <= 0 {
		level = 1
	} else {
		level = level + 1
	}
	keysPerFix, count := ScanRedis(client.Client, match, level)
	out := make(map[string]int, len(keysPerFix))
	for _, v := range keysPerFix {
		out[v] += 1
	}
	res := make(map[string]interface{})
	res["keys"] = out
	res["count"] = count
	res["child"] = false
	if count == 1 && strings.Contains(keysPerFix[0], "*") {
		res["child"] = true
	}
	if count > 1 {
		res["child"] = true
	}
	res["level"] = level
	return res, nil
}

func ScanRedis(client *trace_redis.RedisClient, match string, level int) ([]string, int) {
	var cursor uint64
	var n int
	var totalKeys []string
	level += 1
	for {
		var keys []string
		var err error
		//*扫描所有key，每次20条
		keys, cursor, err = client.Scan(cursor, match, 30).Result()
		if err != nil {
			break
		}
		n += len(keys)
		for _, v := range keys {
			vl := strings.Split(v, ":")
			vll := len(vl)
			str := ""
			for kk, vv := range vl {
				if kk <= level {
					if str == "" {
						str += vv
					} else {
						str += ":" + vv
					}
				}
			}
			if (vll - 1) > level {
				str += ":*"
			}
			totalKeys = append(totalKeys, str)
		}
		if cursor == 0 {
			break
		}
	}
	return totalKeys, n
}
