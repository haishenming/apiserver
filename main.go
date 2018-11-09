package main

import (
	"apiserver/config"
	"apiserver/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
	
	"github.com/spf13/pflag"

)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	
	pflag.Parse()
	
	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
		
	}
	
	// 设置runmode
	gin.SetMode(viper.GetString("runmode"))
	
	// 创建gin实例
	g := gin.New()
	
	middlewares := []gin.HandlerFunc{}
	
	router.Load(
		g,
		middlewares...
		)
	
	
	go func () {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Println("The router has been deployed successfully.")
	}()
	
	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// 检查是否正常启动
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		
		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

