package storage

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-api/config"
	"go-api/utils"
	"log"
	"path"
	"strconv"
	"time"
)

func Upload(localFileName string, cdnPrefix string) (string, error) {

	conf := config.GetConfig()
	ossMap := conf.GetStringMap("oss")
	endpoint := ossMap["endpoint"].(string)
	accessKeyId := ossMap["access_key_id"].(string)
	accessKeySecret := ossMap["access_key_secret"].(string)
	bucketName := ossMap["bucket"].(string)

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)

	if err != nil {
		log.Println(err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Println(err)
	}

	_, filename := path.Split(localFileName)
	hashRandom := utils.GetHash(6)
	hashValue := utils.GetHash(6)
	hashDir := hashRandom[:2] + "/" + hashRandom[2:4] + "/" + hashRandom[4:]
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	objectName := cdnPrefix + "/" + hashDir + timestamp + "_" + hashValue + "_" + filename

	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		log.Println(err)
	}

	//err = os.Remove(localFileName)

	if err != nil {
		return "", err
	}
	return objectName, nil
}

func DeleteOss(cdnFilename string) error {

	if cdnFilename == "" {
		return nil
	}
	conf := config.GetConfig()
	ossMap := conf.GetStringMap("oss")
	endpoint := ossMap["endpoint"].(string)
	accessKeyId := ossMap["access_key_id"].(string)
	accessKeySecret := ossMap["access_key_secret"].(string)
	bucketName := ossMap["bucket"].(string)

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)

	if err != nil {
		log.Println(err)
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Println(err)
		return err
	}

	err = bucket.DeleteObject(cdnFilename[1:])
	if err != nil {
		log.Println(err)
	}
	return err
}
