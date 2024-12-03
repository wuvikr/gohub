package sms

import (
	"encoding/json"
	"gohub/pkg/config"
	"gohub/pkg/logger"

	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Aliyun struct{}

func (a *Aliyun) Send(phone string, message Message) bool {
	// 从配置中获取阿里云短信配置
	aliyunConfig := config.GetStringMapString("sms.aliyun")

	// 打印配置信息用于调试
	logger.DebugJSON("短信[阿里云]", "配置信息", aliyunConfig)

	if len(aliyunConfig) == 0 {
		logger.ErrorString("短信[阿里云]", "配置错误", "找不到阿里云短信配置信息")
		return false
	}

	// 检查必要的配置项
	requiredConfigs := []string{"access_key_id", "access_key_secret", "sign_name"}
	for _, c := range requiredConfigs {
		if aliyunConfig[c] == "" {
			logger.ErrorString("短信[阿里云]", "配置错误", c+" 未配置")
			return false
		}
	}

	clientConfig := &openapi.Config{
		AccessKeyId:     tea.String(aliyunConfig["access_key_id"]),
		AccessKeySecret: tea.String(aliyunConfig["access_key_secret"]),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
		RegionId:        tea.String("cn-hangzhou"), // 添加区域配置
	}

	// 打印客户端配置用于调试
	logger.DebugJSON("短信[阿里云]", "客户端配置", map[string]string{
		"AccessKeyId": *clientConfig.AccessKeyId,
		"Endpoint":    *clientConfig.Endpoint,
		"RegionId":    *clientConfig.RegionId,
	})

	smsClient, err := dysmsapi.NewClient(clientConfig)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "创建短信客户端失败", err.Error())
		return false
	}

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "模板参数序列化失败", err.Error())
		return false
	}

	// 打印请求参数用于调试
	logger.DebugJSON("短信[阿里云]", "请求参数", map[string]string{
		"PhoneNumbers":  phone,
		"SignName":      aliyunConfig["sign_name"],
		"TemplateCode":  message.Template,
		"TemplateParam": string(templateParam),
	})

	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(aliyunConfig["sign_name"]),
		TemplateCode:  tea.String(message.Template),
		TemplateParam: tea.String(string(templateParam)),
	}

	// 添加更多的运行时选项
	runtime := &util.RuntimeOptions{
		ReadTimeout:    tea.Int(5000),
		ConnectTimeout: tea.Int(5000),
	}

	response, err := smsClient.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "发送短信失败", err.Error())
		return false
	}

	if response.Body == nil || *response.Body.Code != "OK" {
		if response.Body != nil {
			logger.ErrorString("短信[阿里云]", "服务商返回错误", *response.Body.Message)
		} else {
			logger.ErrorString("短信[阿里云]", "服务商返回错误", "response body is nil")
		}
		return false
	}

	logger.DebugString("短信[阿里云]", "发送成功", "")
	return true
}
