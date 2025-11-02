package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Port int
	DB   *DBConfig
}

type DBConfig struct {
	User          string
	Password      string
	Host          string
	Port          int
	Name          string
	EnableSSLMode bool
}

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load config from env: ", err)
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == " " {
		fmt.Println("Failed to load http port from env")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must me a number")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB Host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB port is required")
		os.Exit(1)
	}

	dbPrt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("DB Port must be a number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB User is required")
		os.Exit(1)
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		fmt.Println("DB Password is required")
		os.Exit(1)
	}

	enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")

	enableSSLMode, err := strconv.ParseBool(enableSslMode)
	if err != nil {
		fmt.Println("Invalid enable ssl mode value", err)
		os.Exit(1)
	}
	dbConfig := &DBConfig{
		User:          dbUser,
		Password:      dbPass,
		Host:          dbHost,
		Port:          int(dbPrt),
		Name:          dbName,
		EnableSSLMode: enableSSLMode,
	}

	configuration = &Config{
		Port: int(port),
		DB:   dbConfig,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}

	return configuration
}
