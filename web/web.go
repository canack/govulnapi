package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed website/*
var websiteDir embed.FS

type Web struct {
	router        *chi.Mux
	listenAddress string
	fs            http.FileSystem
}

func New(listenAddress string) *Web {
	sub, _ := fs.Sub(websiteDir, "website")
	web := Web{
		router:        chi.NewRouter(),
		listenAddress: listenAddress,
		fs:            http.FS(sub),
	}
	return &web
}

func (web *Web) Run() {
	log.Println("Starting web client ...")
	web.setupRoutes()
	log.Fatalln(http.ListenAndServe(web.listenAddress, web.router))
}

func (web *Web) setupRoutes() {
	web.router.Get("/*", web.handleWebsite)
}

func (web *Web) handleWebsite(w http.ResponseWriter, r *http.Request) {
	http.FileServer(web.fs).ServeHTTP(w, r)
}
