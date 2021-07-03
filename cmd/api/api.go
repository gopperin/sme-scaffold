package api

import (
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	myrouter "sme-stage/router"
)

// StartCmd api
var (
	StartCmd = &cobra.Command{
		Use:   "api",
		Short: "start sme-stage api",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			logrus.Println("sme-stage start")

			//启动API服务
			startAPI()

			logrus.Println("sme-stage end")
		},
	}
)

func startAPI() {
	router := gin.Default()

	router.Use(Cors())

	/* api base */
	myrouter.SetupBaseRouter(router)

	server := &http.Server{
		Addr:         ":" + "12921",
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	gracehttp.Serve(server)
}
