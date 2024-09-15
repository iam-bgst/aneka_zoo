package config

import (
	"gorm.io/gorm"
	sql "zoo/datasource"
)

type (
	PostgresGorm *gorm.DB
)

func (db *Database) NewDatabase() *PostgresGorm {
	postgresGorm := sql.NewDatabase(sql.DatasourcePostgres{
		Dsn:                db.Uri,
		MigrationDirectory: db.MigrationsDir,
	})
	
	return (*PostgresGorm)(&postgresGorm)
}
