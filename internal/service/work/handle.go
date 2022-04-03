package work

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fighthorse/redisAdmin/component/thirdpart/trace_redis"
	"github.com/fighthorse/redisAdmin/internal/pkg/redis"
	"github.com/fighthorse/redisAdmin/protos"
	"github.com/gin-gonic/gin"

	goredis "github.com/go-redis/redis"
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

func HandleErrMsg(title string, err error) string {
	if err != nil {
		return fmt.Sprintf("%s,[%s]", title, err.Error())
	}
	return title
}

func HandleByType(c *gin.Context, req protos.SearchKeyReq, client *trace_redis.RedisClient) (interface{}, error) {
	ctx := c.Request.Context()
	out := protos.KeysInfo{
		Keys: req.Key,
	}
	var ss string
	var err error
	switch req.Type {
	case "DEL":
		ss, err := client.Del(ctx, req.Key)
		out.Data = HandleErrMsg(fmt.Sprintf("删除状态:%d", ss), err)
		out.Type = "msg"
		return out, err
	case "GET":
		ss, err := client.Get(ctx, req.Key)
		out.Value = ss
		out.Type = "string"
		ttl, _ := client.TTL(ctx, req.Key)
		out.Ttl = ttl
		return out, err
	case "SET":
		ttl := 0
		if req.Ttl < 0 {
			ttl = -1
		} else {
			ttl = req.Ttl
		}
		ss, err := client.Set(ctx, req.Key, req.Value, time.Duration(ttl)*time.Second)
		out.Data = HandleErrMsg(fmt.Sprintf("设置状态:%s", ss), err)
		out.Type = "msg"
		return out, err
	case "HSET":
		values := strings.Split(req.Value, ",")
		valueM := make(map[string]interface{})
		for _, vv := range values {
			if vv == "" {
				continue
			}
			ll := strings.Split(vv, " ")
			if len(ll) >= 1 {
				valueM[ll[0]] = ll[1]
			}
		}

		if len(valueM) > 0 {
			ss, err = client.HMSet(ctx, req.Key, valueM)
		} else {
			ss = "数据不正确"
		}
		out.Data = HandleErrMsg(fmt.Sprintf("设置状态:%s", ss), err)
		out.Type = "msg"
		return out, err
	case "HGET":
		return GetKeyByType(c, req, "hash", req.Key, client), nil
	case "SADD":
		vals := strings.Split(req.Value, ",")
		ssI, err := client.SAdd(ctx, req.Key, vals)
		out.Data = HandleErrMsg(fmt.Sprintf("添加集合数据:%d个", ssI), err)
		out.Type = "msg"
		return out, err
	case "SREM":
		vals := strings.Split(req.Value, ",")
		ssI, err := client.SRem(ctx, req.Key, vals)
		out.Data = HandleErrMsg(fmt.Sprintf("删除集合数据:%d个", ssI), err)
		out.Type = "msg"
		return out, err
	case "SMEMBERS":
		return GetKeyByType(c, req, "set", req.Key, client), nil
	case "ZADD":
		ls := strings.Split(req.Value, " ")
		if len(ls) > 1 {
			score, _ := strconv.ParseFloat(ls[0], 64)
			item := goredis.Z{
				Score:  score,
				Member: ls[1],
			}
			_, _ = client.ZAdd(ctx, req.Key, item)
		}
		return GetKeyByType(c, req, "zset", req.Key, client), nil

	case "ZREM":
		ssI, err := client.ZRem(ctx, req.Key, req.Value)
		out.Data = HandleErrMsg(fmt.Sprintf("删除集合数据:%d个", ssI), err)
		out.Type = "msg"
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
		res.Length = hlen
		res.Value = fmt.Sprintf("%d", hlen)
		if hlen > DataPageSize {
			out := make(map[string]string)
			var cursor uint64
			var keyRes []string
			var page int
			pre := ""
			for {
				var err error
				if page > req.Page {
					break
				}
				//*扫描所有key，每次20条
				keyRes, cursor, err = client.HScan(keys, cursor, "", DataPageSize).Result()
				if err != nil {
					break
				}
				if page == req.Page {
					for k, v := range keyRes {
						if k%2 != 0 {
							out[pre] = v
							continue
						}
						pre = v
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
		setLen, _ := client.SCard(ctx, keys)
		res.Length = setLen
		out := make(map[string]string)
		if setLen > DataPageSize {
			var cursor uint64
			var keyRes []string
			var page int
			for {
				var err error
				if page > req.Page {
					break
				}
				//*扫描所有key，每次20条
				keyRes, cursor, err = client.SScan(keys, cursor, "", DataPageSize).Result()
				if err != nil {
					break
				}
				if page == req.Page {
					for _, v := range keyRes {
						out[v] = v
					}
					break
				}
				page += 1
				if cursor == 0 {
					break
				}
			}
			res.Hash = out
			if len(out) > int(DataPageSize) {
				res.Total = 0
			} else {
				res.Total = int(setLen / DataPageSize)
			}

		} else {
			result, _ := client.SMembers(ctx, keys)
			for _, v := range result {
				out[v] = v
			}
			res.Hash = out
			res.Total = 0
		}
	case "zset":
		setLen, _ := client.ZCard(ctx, keys)
		res.Length = setLen
		var newout []protos.ZSET
		if setLen > DataPageSize {
			var cursor uint64
			var keyRes []string
			var page int = 0
			var pre = ""
			for {
				var err error
				if page > req.Page {
					break
				}
				//*扫描所有key，每次20条
				keyRes, cursor, err = client.ZScan(keys, cursor, "", DataPageSize).Result()
				if err != nil {
					break
				}
				if page == req.Page {
					for k, v := range keyRes {
						if k > 0 && k%2 != 0 {
							score, _ := strconv.ParseFloat(v, 64)
							item := protos.ZSET{
								Score:  score,
								Member: pre,
							}
							newout = append(newout, item)
							continue
						}
						pre = v
					}
					break
				}
				page += 1
				if cursor == 0 {
					break
				}
			}
			res.Zset = newout
			if len(newout) > int(DataPageSize) {
				res.Total = 0
			} else {
				res.Total = int(setLen / DataPageSize)
			}
		} else {
			result, _ := client.ZRangeWithScores(ctx, keys, 0, -1)
			for _, v := range result {
				item := protos.ZSET{
					Score:  v.Score,
					Member: v.Member,
				}
				newout = append(newout, item)
			}
			res.Zset = newout
			res.Total = 0
		}

	case "list":
		listLen, _ := client.LLen(ctx, keys)
		res.Length = listLen
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
