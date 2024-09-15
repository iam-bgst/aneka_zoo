package config

import (
	"log"
	"os"
	
	"github.com/joho/godotenv"
)

type (
	Configurations struct {
		Info      *AppInfo
		Const     *Const
		Databases *Databases
	}
	AppInfo struct {
		AppName    string
		AppVersion string
	}
	Const struct {
		HtppPort string
	}
	Databases struct {
		Postgres *Database
	}
	Database struct {
		Uri           string
		MigrationsDir string
	}
)

func NewConfigurations() *Configurations {
	err := godotenv.Load()
	if err != nil {
		log.Println("Load config from os environment", err)
	}
	return &Configurations{
		Info:      loadAppInfo(),
		Const:     loadConst(),
		Databases: loadDatabases(),
	}
}

func loadAppInfo() *AppInfo {
	return &AppInfo{
		AppName:    os.Getenv("APP_NAME"),
		AppVersion: os.Getenv("APP_VERSION"),
	}
}

func loadConst() *Const {
	return &Const{
		HtppPort: os.Getenv("HTTP_PORT"),
	}
}

func loadDatabases() *Databases {
	return &Databases{
		Postgres: &Database{
			Uri:           os.Getenv("POSTGRES_URI"),
			MigrationsDir: os.Getenv("POSTGRES_MIGRATIONS_DIR"),
		},
	}
}
