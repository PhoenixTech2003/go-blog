package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag" // swagger general

	docs "github.com/PhoenixTech2003/go-blog/docs"
	"github.com/PhoenixTech2003/go-blog/internal/api"
	"github.com/PhoenixTech2003/go-blog/internal/database"
	"github.com/PhoenixTech2003/go-blog/internal/web"
)

//go:embed internal/web/static/*
var staticFiles embed.FS

// @title go-blog API
// @version 1.0
// @description Simple http server for creating go-blog and hosting go-blog
func main() {
	_ = godotenv.Load()

	port := getenvDefault("PORT", "80")
	addr := ":" + port
	baseURL := publicBaseURL(port)
	swaggerHost := publicSwaggerHost(port)

	docs.SwaggerInfo.Host = swaggerHost

	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", baseURL)),
	))
	cssFS, _ := fs.Sub(staticFiles, "internal/web/static/css")
	jsFS, _ := fs.Sub(staticFiles, "internal/web/static/js")

	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		slog.Error("an error occured while opening the database connection", "error", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			slog.Error("an error occured while closing the database connection", "error", err)
		}
	}()

	dbConn := database.New(db)
	api.ApiConfig.Db = dbConn
	web.WebConfig.Db = dbConn

	mux.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.FS(cssFS))))
	mux.Handle("/static/js/", http.StripPrefix("/static/js", http.FileServer(http.FS(jsFS))))

	mux.Handle("/v1/api/articles/", http.StripPrefix("/v1/api/articles", api.ArticlesRouter()))
	mux.Handle("/v1/api/tags/", http.StripPrefix("/v1/api/tags", api.TagsRouter()))

	mux.Handle("/compose/", http.StripPrefix("/compose", web.ComposeRouter()))
	mux.Handle("/", web.IndexRouter())

	server := http.Server{
		Handler:           mux,
		Addr:              addr,
		ReadHeaderTimeout: 5 * time.Second,
	}

	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("server is listening", "addr", addr)
		slog.Info("Swagger UI", "url", fmt.Sprintf("%s/swagger/index.html", baseURL))
		serverErrors <- server.ListenAndServe()
	}()

	shutdownSignal, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %q", err)
		}
	case <-shutdownSignal.Done():
		slog.Info("shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Failed to shut down server: %q", err)
		}

		if err := <-serverErrors; err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to stop server: %q", err)
		}
	}

	slog.Info("server stopped")
}

func getenvDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func publicBaseURL(port string) string {
	if value := os.Getenv("BASE_URL"); value != "" {
		return strings.TrimRight(value, "/")
	}
	if value := os.Getenv("VERCEL_URL"); value != "" {
		return "https://" + strings.TrimRight(value, "/")
	}
	return fmt.Sprintf("http://localhost:%s", port)
}

func publicSwaggerHost(port string) string {
	if value := os.Getenv("SWAGGER_HOST"); value != "" {
		return value
	}
	if value := os.Getenv("VERCEL_URL"); value != "" {
		return value
	}
	return fmt.Sprintf("localhost:%s", port)
}
