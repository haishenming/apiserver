package main

import (
	"apiserver/config"
	"apiserver/model"
	"apiserver/router"
	"apiserver/router/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"
	"time"

	"github.com/spf13/pflag"

	"github.com/gin-contrib/pprof"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.

// @contact.name API Support
func main() {

	pflag.Parse()

	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)

	}

	// 设置runmode
	gin.SetMode(viper.GetString("runmode"))

	// 初始化DB
	model.DB.Init()
	defer model.DB.Close()

	// 创建gin实例
	g := gin.New()

	// pprof router
	pprof.Register(g)

	middlewares := []gin.HandlerFunc{
		middleware.RequestID(),
		middleware.Logging(),
	}

	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// 启动https
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Infof(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
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
		log.Infof("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
