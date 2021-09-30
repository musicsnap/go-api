package main


import (
	"flag"
	"fmt"
	"log"
	"os"

	"go-api/config"
	"go-api/routers"
	"go-api/storage"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	fmt.Println("Current Env: ", *environment)
	config.Init(*environment)

	//2、db存储
	storage.InitDB()
	storage.InitRedis()
	storage.InitMongo()
	storage.InitElasticsearch()
	storage.InitMns()

	err := os.MkdirAll("./locks", 0777)
	if err != nil {
		log.Println("cannot mkdir")
		os.Exit(1)
	}
	//3、路由
	conf := config.GetConfig()
	r := routers.NewRouter()
	err = r.Run(":" + conf.GetString("server.port"))
	if err != nil {
		log.Println("server start failed")
		os.Exit(1)
	}
}
