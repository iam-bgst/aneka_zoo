package datasource

import (
	"database/sql"
	"gorm.io/gorm/logger"
	"log"
	"time"
	
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatasourcePostgres struct {
	Dsn                string
	MigrationDirectory string
}

func NewDatabase(cfg DatasourcePostgres) *gorm.DB {
	sqlDB, err := sql.Open("postgres", cfg.Dsn)
	if err != nil {
		log.Panic("Error occurred while connecting with the database", err)
	}
	
	// Create the connection pool
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//db, err := gorm.Open(postgres.New(postgres.Config{
	//	Conn: sqlDB,
	//}), &gorm.Config{})
	//
	// ping
	if err = sqlDB.Ping(); err != nil {
		log.Panic("Error occurred while connecting with the database")
	}
	
	return db
}
