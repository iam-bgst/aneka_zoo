package application

import (
	"context"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	
	"zoo/config"
	"zoo/libraries/logger"
)

type (
	App struct {
		TerminalHandler chan os.Signal
		Context         context.Context
		Logger          logger.ILogger
		Configurations  *config.Configurations
		Postgres        *config.PostgresGorm
	}
)

func NewApp(ctx context.Context) (*App, error) {
	configuration := config.NewConfigurations()
	
	terminalHandler := make(chan os.Signal, 1)
	signal.Notify(terminalHandler,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	
	// init datasource
	postgres := configuration.Databases.Postgres.NewDatabase()
	
	// init logger
	loggerLib := logger.NewLogger(&logger.Option{
		Level:     "info",
		Out:       os.Stdout,
		AppName:   configuration.Info.AppName,
		Formatter: logger.FormatJSON,
	})
	
	return &App{
		TerminalHandler: terminalHandler,
		Context:         ctx,
		Logger:          loggerLib,
		Configurations:  configuration,
		Postgres:        postgres,
	}, nil
	
}

func (app *App) Close(ctx context.Context) {
	if app.Postgres != nil {
		sqlDB, _ := (*gorm.DB)(*app.Postgres).DB()
		err := sqlDB.Close()
		if err != nil {
			return
		}
	}
	app.Logger.InfoWithContext(ctx, "", "APP SUCCESSFULLY CLOSED")
}
