package httpserver

import "github.com/fighthorse/redisAdmin/component/httpclient"

var (
	Amap = &AmapClient{}
)

func Init() {

	Amap.name = "amap"
	Amap.Client = httpclient.New(Amap.name)

}
