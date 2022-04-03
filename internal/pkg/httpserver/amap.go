package httpserver

import (
	"context"

	"github.com/fighthorse/redisAdmin/component/httpclient"
	"github.com/fighthorse/redisAdmin/component/log"
	"github.com/fighthorse/redisAdmin/protos"
)

// 高德地图 api
// https://lbs.amap.com/api/webservice/guide/api/weatherinfo

type AmapClient struct {
	*httpclient.Client
	name string
}

func (c *AmapClient) GetWeatherInfo(ctx context.Context, reqBody map[string]interface{}) (result protos.WeatherInfoRep, err error) {

	err = c.Client.Send(ctx, "GET", "/ip/grade/get", reqBody, &result)
	if err != nil {
		log.Error(ctx, "GetWeatherInfoErr", log.Fields{"in": reqBody, "out": result, "err": err.Error()})
	}
	return result, err
}
