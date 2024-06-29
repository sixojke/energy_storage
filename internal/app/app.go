package app

import (
	"context"
	"database/sql"
	"energy_storage/internal/config"
	"energy_storage/internal/devilery"
	"energy_storage/internal/repository"
	"energy_storage/internal/server"
	"energy_storage/internal/service"
	"energy_storage/pkg/db"
	"energy_storage/pkg/logger"
	"energy_storage/pkg/migrate"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("[CONFIG] Config successfully initialized")

	postgres, err := db.NewPostgresDB(db.PostgresConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		logger.Fatal(err)
	}
	defer postgres.Close()
	logger.Info("[POSTGRES] Connected to postgres")

	migratorPostgres, err := MigratorPostgres(postgres.DB)
	if err != nil {
		logger.Fatal(err)
	}

	if err := migratorPostgres.Down(); err != nil {
		logger.Fatal(err)
	}

	if err := migratorPostgres.Up(); err != nil {
		logger.Fatal(err)
	}
	logger.Info("[MIGRATE] Migrations applied successfully")

	repo := repository.NewRepository(&repository.Deps{Postgres: postgres})
	service := service.NewSerivice(&service.Deps{Repo: *repo})
	handler := devilery.NewHandler(service)

	srv := server.NewServer(cfg.Server, handler.Init())
	go func() {
		if err := srv.Start(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %v\n", err)
		}
	}()
	logger.Info(fmt.Sprintf("[SERVER] Started :%v", cfg.Server.Port))

	shutdown(srv, postgres)
}

func MigratorPostgres(postgres *sql.DB) (*migrate.Migrate, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %v", err)
	}

	schemaPath := "file://" + path.Join(dir, "schema", "postgres")

	return migrate.NewMigratorPostgres(schemaPath, postgres)
}

func shutdown(srv *server.Server, postgres *sqlx.DB) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 3 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	postgres.Close()
}
