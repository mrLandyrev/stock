package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/mrLandyrev/stock/internal/app/rest"
	"github.com/mrLandyrev/stock/internal/respositories/placements"
	"github.com/mrLandyrev/stock/internal/uescases/get_balance"
	"github.com/mrLandyrev/stock/internal/uescases/reserve"
	"github.com/mrLandyrev/stock/internal/uescases/unreserve"
	"go.uber.org/zap"
)

type AppConfig struct {
	DatabaseDriver     string
	DatabaseConnection string
	Address            string
}

type App struct {
	server *http.Server
	db     *sql.DB
	logger *zap.Logger
}

func NewApp(config AppConfig) *App {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("Error at create logger")
	}

	db, err := sql.Open(config.DatabaseDriver, config.DatabaseConnection)
	if err != nil {
		logger.Fatal("Error at connect to database")
	}

	placementsRepository := placements.NewPlacementsRepository(db)
	reserveUseCase := reserve.NewReserveUseCase(placementsRepository)
	unreserveUseCase := unreserve.NewUnreserveUseCase(placementsRepository)
	getBalanceUseCase := get_balance.NewGetBalanceUseCase(placementsRepository)

	server := &http.Server{
		Addr: config.Address,
		Handler: rest.BuildHandlers(
			getBalanceUseCase,
			reserveUseCase,
			unreserveUseCase,
		),
	}

	return &App{
		db:     db,
		server: server,
		logger: logger,
	}
}

func (app *App) Run() {
	app.logger.Info("Run application")
	go func() {
		app.server.ListenAndServe()
	}()
}

func (app *App) Stop() {
	app.logger.Info("Gracefull stop")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.server.Shutdown(ctx)
	app.db.Close()
	app.logger.Info("Stoped")
}
