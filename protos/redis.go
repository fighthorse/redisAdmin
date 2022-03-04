package protos

import "time"

type AddCfgReq struct {
	Name  string `form:"name" json:"name" mapstructure:"name"`
	Addr  string `form:"addr" json:"addr" mapstructure:"addr"`
	Pwd   string `form:"pwd" json:"pwd" mapstructure:"pwd"`
	Token string `form:"token" json:"token" mapstructure:"token"`
}

type SearchReq struct {
	Client string `form:"client" json:"client" mapstructure:"client"`
	Db     string `form:"db" json:"db" mapstructure:"db"`
	Key    string `form:"key" json:"key" mapstructure:"key"`
	Token  string `form:"token" json:"token" mapstructure:"token"`
	Level  int    `form:"level" json:"level" mapstructure:"level"`
}

type SearchKeyReq struct {
	Client string `form:"client" json:"client" mapstructure:"client"`
	Db     string `form:"db" json:"db" mapstructure:"db"`
	Type   string `form:"type" json:"type" mapstructure:"type"`
	Key    string `form:"key" json:"key" mapstructure:"key"`
	Value  string `form:"value" json:"value" mapstructure:"value"`
	Ttl    int    `form:"ttl" json:"ttl" mapstructure:"ttl"`
	Token  string `form:"token" json:"token" mapstructure:"token"`
	Level  int    `form:"level" json:"level" mapstructure:"level"`
	Page   int    `form:"page" json:"page" mapstructure:"page"`
}

type KeysInfo struct {
	Type  string            `form:"type" json:"type" mapstructure:"type"`
	Keys  string            `form:"keys" json:"keys" mapstructure:"keys"`
	Value string            `form:"value" json:"value" mapstructure:"value"`
	Ttl   time.Duration     `form:"ttl" json:"ttl" mapstructure:"ttl"`
	List  []ListRes         `form:"list" json:"list" mapstructure:"list"`
	Hash  map[string]string `form:"hash" json:"hash" mapstructure:"hash"`
	Page  int               `form:"page" json:"page" mapstructure:"page"`
	Total int               `form:"total" json:"total" mapstructure:"total"`
}

type ListRes struct {
	Index int    `form:"index" json:"index" mapstructure:"index"`
	Value string `form:"value" json:"value" mapstructure:"value"`
}
