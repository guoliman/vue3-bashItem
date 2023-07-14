package email

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"vue3-bashItem/pkg/settings"
)

// EmailContent 发送邮件所需参数
type EmailContent struct {
	Token string `json:"token"`
	Contacts []string `json:"contacts"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

// RespItemData 返回结果
type RespItemData struct {
	User string `json:"user"`
	Status string `json:"status"`
	Message string `json:"message"`
}

type Resp struct {
	Code int
	Data []RespItemData
}

func SendEmail(emailContent EmailContent) (Resp,error) {

	var resp Resp

	client := resty.New()

	response, err1 := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(emailContent).
		Post(settings.EmailSetting.HttpApi)

	err2 := json.Unmarshal([]byte(response.String()), &resp)

	if err1 != nil {
		return resp,err2
	}

	if err2 != nil {
		return resp,err2
	}

	return resp,nil
}