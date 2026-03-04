package main

import (
	"log"
	"net/http"
	"os"

	"u-short/internal/config"
	"u-short/internal/db"

	"u-short/internal/handler"
	"u-short/internal/repository"
	"u-short/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0755)
	}

	dbConn := db.InitDB(cfg.DbUrl)

	repo := repository.NewUrlRepository(dbConn)
	svc := service.NewUrlService(repo)
	hdl := handler.NewUrlHandler(svc)

	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", hdl.Index)
	r.Get("/partials/form-shorten", hdl.ShortLink)
	r.Get("/partials/form-qr", hdl.QRCode)
	r.Post("/shorten", hdl.Create(cfg.BaseUrl))
	r.Post("/scan", hdl.Scan)
	r.Get("/{shortCode}", hdl.Redirect)

	log.Println("Server u-short running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
