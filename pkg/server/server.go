package server

import (
	"fmt"

	"github.com/andrasbarabas/shapeshiftr-api/config"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/l"
	"github.com/andrasbarabas/shapeshiftr-api/routes"
	"github.com/gin-gonic/gin"
)

func Setup() {
	setMode(config.ServerConfig.GinMode)

	// Setup logger
	l.Setup()

	// Setup router
	router := routes.Setup()

	startServer(router)
}

func setMode(m string) {
	gin.SetMode(m)
}

func startServer(r *gin.Engine) {
	host := config.ServerConfig.ApplicationHost
	port := config.ServerConfig.ApplicationPort
	address := fmt.Sprintf("%v:%v", host, port)

	err := r.Run(address)

	if err != nil {
		fmt.Println(err)
	}
}
