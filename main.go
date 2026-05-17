package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
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

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8081"
	}
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:8081"
	}

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
	dbConn := database.New(db)
	api.ApiConfig.Db = dbConn
	web.WebConfig.Db = dbConn

	mux.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.FS(cssFS))))
	mux.Handle("/static/js/", http.StripPrefix("/static/js", http.FileServer(http.FS(jsFS))))

	apiPrefix := "/v1/api/go-blog"
	mux.Handle(apiPrefix+"/", http.StripPrefix(apiPrefix, api.QouteRouter()))

	mux.Handle("/compose/", http.StripPrefix("/compose", web.ComposeRouter()))
	mux.Handle("/", web.IndexRouter())

	addr := ":8081"
	server := http.Server{
		Handler:           mux,
		Addr:              addr,
		ReadHeaderTimeout: 5 * time.Second,
	}
	slog.Info("server is listening", "addr", addr)
	slog.Info("Swagger UI", "url", fmt.Sprintf("%s/swagger/index.html", baseURL))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %q", err)
	}
}
