package db

import (
	"fmt"

	"github.com/andrasbarabas/shapeshiftr-api/model"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/connector"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/l"
	"gorm.io/gorm"
)

var databaseInstance *gorm.DB

func GetConnection() *gorm.DB {
	return databaseInstance
}

func Setup(c *connector.ConnectorConfiguration) {
	databaseInstance = connector.GetDatabaseInstance(c)

	migrate()
}

func migrate() {
	err := databaseInstance.Table("food").AutoMigrate(&model.Food{})

	if err != nil {
		l.Logger.Fatal(err.Error())

		return
	}

	fmt.Println("Migration successfully completed!")
}
