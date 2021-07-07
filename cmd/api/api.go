package api

import (
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	myrouter "sme-stage/router"
	myconfig "sme-stage/utils/config"
)

// StartCmd api
var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "start sme-stage api", SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			//启动API服务
			run()

			logrus.Println("sme-stage end")
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

	/* api base */
	myrouter.SetupBaseRouter(router)

	server := &http.Server{
		Addr:         ":" + myconfig.Application.Port,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	logrus.Println("sme-stage start on:", myconfig.Application.Port)
	gracehttp.Serve(server)
}
