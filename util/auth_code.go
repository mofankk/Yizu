package yizuutil

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func SendAuthCode(phoneNum, captcha string) bool {

	credential := common.NewCredential(
		"AKIDrTosXySFc8yb9UjcOMUOt86c8agkmF8c",
		"54nGE4pJo3INFWJPKL6rQmV3Qd4SgS5T",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-nanjing", cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + phoneNum})
	request.SmsSdkAppId = common.StringPtr("1400525279")
	request.SignName = common.StringPtr("崔有朋我的技术测试")
	request.TemplateId = common.StringPtr("967438")
	request.TemplateParamSet = common.StringPtrs([]string{captcha, "6"})

	_, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false
	}
	if err != nil {
		log.Errorf("验证码发送失败: %v", err)
		return false
	}
	return true
}
