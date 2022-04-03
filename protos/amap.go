package protos

type WeatherInfoRep struct {
	Status string      `json:"status"` // 1：成功；0：失败
	Count  string      `json:"count"`  // 返回结果总数目
	Info   string      `json:"info"`   // 返回的状态信息
	Lives  []LivesInfo `json:"lives"`
}

type LivesInfo struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	Reporttime    string `json:"reporttime"`
}
