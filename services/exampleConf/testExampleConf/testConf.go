package testExampleConf

import "github.com/google/uuid"

// 200 例子 如下2个是组合
type Data struct {
	ID   int       `json:"id" format:"int64" example:"1"`
	Name string    `json:"name" example:"gavin" `
	Text string    `json:"title" example:"Object data"`
	UUID uuid.UUID `json:"uuid" format:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
}
type Response struct {
	Title      map[string]string      `json:"title" example:"en:Map,ru:Карта,kk:Карталар"`
	CustomType map[string]interface{} `json:"map_data" swaggertype:"object,string" example:"key:value,key2:value2"`
	Object     Data                   `json:"object"`
}

// HTTPSuccess 200 例子
type HTTPSuccess struct {
	Code int               `json:"code" example:"200"`
	Data map[string]string `json:"data" example:"Token:eyJhbGciOiCI6IkpXVCJ9"`
	Msg  string            `json:"msg" example:"null"`
}

// HTTPError 400 例子
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
