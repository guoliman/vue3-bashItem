package dictMethod

// DictData 字段类型
type DictData struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Status     uint8  `json:"status"`
	Remark     string `json:"remark"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// DictSonData 字段子级类型
type DictSonData struct {
	Id         int64  `json:"id"`
	TypeCode   string `json:"type_code"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	Sort       int64  `json:"Sort"`
	Status     uint8  `json:"status"`
	Defaulted  uint8  `json:"defaulted"`
	Remark     string `json:"remark"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// DelId 删除id
type DelId struct {
	Id string `json:"id"`
}
