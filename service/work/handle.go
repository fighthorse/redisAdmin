package work

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/fighthorse/redisAdmin/pkg/redis"
	"github.com/fighthorse/redisAdmin/pkg/thirdpart/trace_redis"
	"github.com/fighthorse/redisAdmin/protos"
	"github.com/gin-gonic/gin"
)

var (
	DataPageSize int64 = 20
)

func HandleKey(c *gin.Context, req protos.SearchKeyReq) (interface{}, error) {
	// get client
	db, _ := strconv.Atoi(req.Db)
	client := redis.LoadOthersDB(req.Client, db)
	if client == nil {
		return nil, errors.New("redis client create error")
	}
	// handleKey
	return HandleByType(c, req, client.Client)
}

func HandleByType(c *gin.Context, req protos.SearchKeyReq, client *trace_redis.RedisClient) (interface{}, error) {
	out := make(map[string]interface{})
	switch req.Type {
	case "GET":
		ss, err := client.Get(c.Request.Context(), req.Key)
		out["data"] = ss
		out["type"] = "string"
		return out, err
	case "SET":
		ttl := 0
		if req.Ttl < 0 {
			ttl = -1
		} else {
			ttl = req.Ttl
		}
		ss, err := client.Set(c.Request.Context(), req.Key, req.Value, time.Duration(ttl)*time.Second)
		out["data"] = ss
		out["type"] = "string"
		return out, err
	}
	return nil, nil
}

func GetKeyByType(c *gin.Context, req protos.SearchKeyReq, typeInfo, keys string, client *trace_redis.RedisClient) protos.KeysInfo {
	ctx := c.Request.Context()
	res := protos.KeysInfo{}
	res.Keys = keys
	res.Type = typeInfo
	if req.Page < 0 {
		req.Page = 0
	}
	res.Page = req.Page

	// ttl 过期时间
	ttl, _ := client.TTL(ctx, keys)
	res.Ttl = ttl

	// 根据类型处理
	value := ""
	switch typeInfo {
	case "string":
		value, _ = client.Get(ctx, keys)
		res.Value = value
	case "hash":
		hlen, _ := client.HLen(ctx, keys)
		res.Value = fmt.Sprintf("%d", hlen)
		if hlen > DataPageSize {
			out := make(map[string]string)
			var cursor uint64
			var keyRes []string
			var page int
			for {
				var err error
				//*扫描所有key，每次20条
				keyRes, cursor, err = client.HScan(keys, cursor, "*", DataPageSize).Result()
				if err != nil {
					break
				}
				if page == req.Page {
					for k, v := range keyRes {
						pre := v
						if k%2 != 0 {
							out[pre] = v
						}
					}
					break
				}
				page += 1
				if cursor == 0 {
					break
				}
			}
			res.Hash = out
			res.Total = int(hlen / DataPageSize)
		} else {
			res.Hash, _ = client.HGetAll(ctx, keys)
			res.Total = 0
		}
	case "set":

	case "zset":

	case "list":
		listLen, _ := client.LLen(ctx, keys)
		if listLen > DataPageSize {
			start := int64(res.Page) * DataPageSize
			result, _ := client.LRange(keys, start, start+DataPageSize).Result()
			for k, v := range result {
				item := protos.ListRes{
					Index: k + int(start),
					Value: v,
				}
				res.List = append(res.List, item)
			}
			res.Total = int(listLen / DataPageSize)
		} else {
			result, _ := client.LRange(keys, 0, -1).Result()
			for k, v := range result {
				item := protos.ListRes{
					Index: k,
					Value: v,
				}
				res.List = append(res.List, item)
			}
			res.Total = 0
		}

	case "bitmap":

	default:
		res.Value = "暂不支持查看类型"
	}
	return res
}
