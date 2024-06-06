package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/godwinrob/harvester/business/sdk/sqldb"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	conf "github.com/ardanlabs/conf/v3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

var build = "develop"

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("startup", "msg", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	slog.Info("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0), "BUILD", build)

	// -------------------------------------------------------------------------
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout        time.Duration `conf:"default:5s"`
			WriteTimeout       time.Duration `conf:"default:10s"`
			IdleTimeout        time.Duration `conf:"default:120s"`
			ShutdownTimeout    time.Duration `conf:"default:20s"`
			APIHost            string        `conf:"default:0.0.0.0:3000"`
			DebugHost          string        `conf:"default:0.0.0.0:3010"`
			CORSAllowedOrigins []string      `conf:"default:*"`
		}
		DB struct {
			User         string `conf:"default:postgres"`
			Password     string `conf:"default:postgres"`
			Host         string `conf:"default:postgres"`
			Name         string `conf:"default:postgres"`
			MaxIdleConns int    `conf:"default:0"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "Harvester",
		},
	}

	const prefix = "HARVESTER"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// App Starting

	slog.Info("starting service", "version", cfg.Build)
	defer slog.Info("shutdown complete")

	slog.Info("startup", "config", fmt.Sprintf("%+v", cfg))

	// -------------------------------------------------------------------------
	// Database Support

	slog.Info("startup", "status", "initializing database support", "hostport", cfg.DB.Host)

	time.Sleep(3 * time.Second)

	db, err := sqldb.Open(sqldb.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping db: %w", err)
	}

	//-------------------------------------------------------------------------
	//Start API Service

	slog.Info("startup", "status", "initializing V1 API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      r,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		slog.Info("startup", "status", "api router started")

		serverErrors <- api.ListenAndServe()
	}()

	//-------------------------------------------------------------------------
	//Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		slog.Info("shutdown", "status", "shutdown started", "signal", sig)
		defer slog.Info("shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
