package storage

import (
	mns "github.com/aliyun/aliyun-mns-go-sdk"
	"go-api/config"
)

type MnsConf struct {
	Url             string `json:"url"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

var AliMnsQueue mns.AliMNSQueue

func InitMns() {

	queueName := "ips-render-gif"
	//queueName := "user-psd-queue-test"

	conf := config.GetConfig()
	mnsConf := conf.GetStringMap("mns")

	Client := mns.NewAliMNSClient(
		mnsConf["url"].(string),
		mnsConf["access_key_id"].(string),
		mnsConf["access_key_secret"].(string),
	)

	AliMnsQueue = mns.NewMNSQueue(queueName, Client)
}

func GetQueue() mns.AliMNSQueue {
	return AliMnsQueue
}
