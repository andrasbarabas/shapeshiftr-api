package l

import (
	"log"

	"github.com/andrasbarabas/shapeshiftr-api/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Setup() {
	var err error

	if config.ServerConfig.GinMode == gin.ReleaseMode {
		Logger, err = zap.NewProduction()
	} else {
		Logger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}

	err = Logger.Sync()

	if err != nil {
		return
	}
}
