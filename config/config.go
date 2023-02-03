package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type app struct {
	ProjectName string
}

type database struct {
	DatabaseDriver   string
	DatabaseHost     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabasePort     string
}

type server struct {
	ApplicationHost string
	ApplicationPort string
	GinMode         string
	TrustedProxies  string
}

// AppConfig with default values
var AppConfig = &app{
	ProjectName: "shapeshiftr-api",
}

// DatabaseConfig with default values
var DatabaseConfig = &database{
	DatabaseDriver:   "postgres",
	DatabaseHost:     "127.0.0.1",
	DatabaseName:     "shapeshiftr-api",
	DatabaseUser:     "shapeshiftr-api",
	DatabasePassword: "s3cr3t",
	DatabasePort:     "5432",
}

// ServerConfig with default values
var ServerConfig = &server{
	ApplicationHost: "0.0.0.0",
	ApplicationPort: "8080",
	GinMode:         gin.DebugMode,
	TrustedProxies:  "127.0.0.1",
}

func Setup() {
	loadEnvironmentVars()

	setupAppConfig()
	setupDatabaseConfig()
	setupServerConfig()
}

func loadEnvironmentVars() {
	err := godotenv.Load()

	if os.Getenv("GIN_MODE") != "release" && err != nil {
		fmt.Printf("Error during loading environment variables: %v\n", err)
	}
}

func setupAppConfig() {
	if os.Getenv("PROJECT_NAME") != "" {
		AppConfig.ProjectName = os.Getenv("PROJECT_NAME")
	}
}

func setupDatabaseConfig() {
	if os.Getenv("DATABASE_DRIVER") != "" {
		DatabaseConfig.DatabaseDriver = os.Getenv("DATABASE_DRIVER")
	}

	if os.Getenv("DATABASE_HOST") != "" {
		DatabaseConfig.DatabaseHost = os.Getenv("DATABASE_HOST")
	}

	if os.Getenv("DATABASE_NAME") != "" {
		DatabaseConfig.DatabaseName = os.Getenv("DATABASE_NAME")
	}

	if os.Getenv("DATABASE_USER") != "" {
		DatabaseConfig.DatabaseUser = os.Getenv("DATABASE_USER")
	}

	if os.Getenv("DATABASE_PASSWORD") != "" {
		DatabaseConfig.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	}

	if os.Getenv("DATABASE_PORT") != "" {
		DatabaseConfig.DatabasePort = os.Getenv("DATABASE_PORT")
	}
}

func setupServerConfig() {
	if os.Getenv("APPLICATION_HOST") != "" {
		ServerConfig.ApplicationHost = os.Getenv("APPLICATION_HOST")
	}

	if os.Getenv("APPLICATION_PORT") != "" {
		ServerConfig.ApplicationPort = os.Getenv("APPLICATION_PORT")
	}

	if os.Getenv("GIN_MODE") != "" {
		ServerConfig.GinMode = os.Getenv("GIN_MODE")
	}

	if os.Getenv("TRUSTED_PROXIES") != "" {
		ServerConfig.TrustedProxies = os.Getenv("TRUSTED_PROXIES")
	}
}
