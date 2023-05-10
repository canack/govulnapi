package main

import (
	"govulnapi/api"
	"govulnapi/coingecko"
	"govulnapi/web"
	"io"
	"log"
	"os"
	"os/signal"
)

//	@title			  Govulnapi
//	@version		  1.0
//	@description	Deliberately vulnerable API written in Go

//	@license.name	MIT
//	@license.url	https://mit-license.org

//	@host		  localhost:8081
//	@BasePath	/api
//	@Schemes	http

//	@securityDefinitions.apikey	Bearer
//	@in	      				   			  header
//	@name						            Authorization
//	@description				        Type "BEARER" followed by a space and the token.

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	// Log to both stdout and file
	// CWE-276: Improper Access Control
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	// Setup servers
	coingecko := coingecko.New(":8082")
	api := api.New(":8081", "http://localhost:8082")
	web := web.New(":8080")

	// Run servers
	go coingecko.Run()
	go api.Run()
	go web.Run()

	// Graceful shutdown for database and logfile
	<-shutdown
	api.Shutdown()
	logFile.Close()
}
