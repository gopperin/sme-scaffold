package api

import (
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	myrouter "sme-scaffold/router"
	myconfig "sme-scaffold/utils/config"
)

// StartCmd api
var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "start sme-scaffold api", SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			//启动API服务
			run()

			logrus.Println("sme-scaffold end")
		},
	}
)

func setup() {

	//1. 读取配置
	myconfig.Setup("./")

}

func run() {

	router := gin.Default()

	router.Use(Cors())

	_db, err := InitDatabase()
	if err != nil {
		logrus.Println("InitDatabase error", err.Error())
		return
	}
	logrus.Info("sme-scaffold start on:", myconfig.Application.Port)

	/* api base */
	myrouter.SetupBaseRouter(router)

	/* product base */
	myrouter.SetupProductRouter(router, _db)

	server := &http.Server{
		Addr:         ":" + myconfig.Application.Port,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	logrus.Println("sme-scaffold start on:", myconfig.Application.Port)
	gracehttp.Serve(server)
}
