package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/mrLandyrev/stock/internal/app"

	"syscall"
)

func buildConfig() app.AppConfig {
	var config app.AppConfig
	config.Address = *flag.String("address", "0.0.0.0:8080", "rest server address")
	config.DatabaseDriver = *flag.String("db-driver", "postgres", "database driver")
	databaseHost := *flag.String("db-host", "db", "database host")
	databasePort := *flag.Int("db-port", 5432, "database port")
	databaseUser := *flag.String("db-user", "postgres", "database user")
	databasePassword := *flag.String("db-password", "local", "database password")
	databaseName := *flag.String("db-name", "postgres", "database name")
	flag.Parse()

	config.DatabaseConnection = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseHost, databasePort, databaseUser, databasePassword, databaseName)

	return config
}

func main() {
	app := app.NewApp(buildConfig())
	app.Run()
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop
	app.Stop()
}
