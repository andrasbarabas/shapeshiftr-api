package main

import (
	"github.com/andrasbarabas/shapeshiftr-api/config"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/connector"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/db"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/server"
)

func main() {
	config.Setup()

	db.Setup(&connector.ConnectorConfiguration{
		DatabaseDriver:   config.DatabaseConfig.DatabaseDriver,
		DatabaseHost:     config.DatabaseConfig.DatabaseHost,
		DatabaseName:     config.DatabaseConfig.DatabaseName,
		DatabasePassword: config.DatabaseConfig.DatabasePassword,
		DatabasePort:     config.DatabaseConfig.DatabasePort,
		DatabaseUser:     config.DatabaseConfig.DatabaseUser,
	})

	server.Setup()
}
