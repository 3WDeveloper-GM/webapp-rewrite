package main

import (
	"flag"
	"net/http"
	"time"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg"
)

func main() {

	addr := flag.String("addr", ":9090", "http network address: A number that defines the addr of our application")
	flag.Parse()

	app := config.ExportLoggers()

	app.Infolog.Printf("Starting application on port %v \n", *addr)

	//err := http.ListenAndServe(addr, Router())
	//if err != nil {
	//	log.Fatal("error on http.ListenAndServe function, check the addr and mux variables.")
	//}
	srv := &http.Server{
		Handler:      Router(),
		Addr:         *addr,
		ErrorLog:     app.Errorlog,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	app.Errorlog.Fatal(srv.ListenAndServe())
}
