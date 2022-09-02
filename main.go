package main

import (
	"flag"
	"fmt"
	"go-api/utils"
	"log"
	"os"
	"runtime"

	"go-api/config"
	"go-api/routers"
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

	num := runtime.NumCPU()
	//在这里设置num-1的cpu运行go程序
	runtime.GOMAXPROCS(num)
	fmt.Println("num =", num)

	fmt.Println(utils.Now())

	//2、db存储
	//bootstrap.InitDB()
	//bootstrap.InitRedis()
	//bootstrap.InitMongo()
	//bootstrap.InitElasticsearch()
	//bootstrap.InitMns()

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
