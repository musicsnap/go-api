package bootstrap

import (
	"fmt"
	"go-api/config"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"time"
)

var EsClient *elastic.Client

func InitElasticsearch() {
	var err error
	conf := config.GetConfig()
	esMap := conf.GetStringMap("elasticsearch")
	url := fmt.Sprintf("http://%v:%v", esMap["host"], esMap["port"])
	EsClient, err = elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(esMap["auth"].(string), esMap["password"].(string)),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(60*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))

	if err != nil {
		panic(err)
	}
}
