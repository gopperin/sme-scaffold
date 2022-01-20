package api

import (
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	myconfig "sme-scaffold/internal/config"
	mymiddleware "sme-scaffold/internal/middleware"
	myrouter "sme-scaffold/internal/router"
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

	router.Use(mymiddleware.Cors())

	db, err := InitDatabase()
	if err != nil {
		logrus.Println("InitDatabase error", err.Error())
		return
	}
	logrus.Info("sme-scaffold start on:", myconfig.Case.Application.Port)

	/* api base */
	myrouter.SetupBaseRouter(router)

	/* product base */
	myrouter.SetupProductRouter(router, db)

	server := &http.Server{
		Addr:         ":" + myconfig.Case.Application.Port,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	logrus.Println("sme-scaffold start on:", myconfig.Case.Application.Port)
	gracehttp.Serve(server)
}
