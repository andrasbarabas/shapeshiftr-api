package connector

import (
	"fmt"

	"github.com/andrasbarabas/shapeshiftr-api/pkg/l"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectorConfiguration struct {
	DatabaseDriver   string
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
}

func GetDatabaseInstance(c *ConnectorConfiguration) *gorm.DB {
	var dsn string
	var err error
	var databaseInstance *gorm.DB

	switch c.DatabaseDriver {
	case "postgres":
		dsn = fmt.Sprintf("%v://%v:%v@%v:%v/%v",
			c.DatabaseDriver,
			c.DatabaseUser,
			c.DatabasePassword,
			c.DatabaseHost,
			c.DatabasePort,
			c.DatabaseName,
		)

		databaseInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	case "mysql":
		l.Logger.Error("Not implemented error")

	case "sqlite":
		l.Logger.Error("Not implemented error")
	}

	if err != nil {
		l.Logger.Error(err.Error())
	}

	return databaseInstance
}
