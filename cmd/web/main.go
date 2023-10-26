package main

import (
	"flag"
	"net/http"
	"time"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg"
)

func main() {

	//setting the flags of the program, i can control the port and the database information, I'm using
	//MariaDB as a database software right now.
	addr := flag.String("addr", ":9090", "http network address: A number that defines the addr of our application")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL dsn")
	flag.Parse()

	app := config.ExportConfig(*dsn)

	app.Infolog.Printf("Starting application on port %v \n", *addr) //Logging on start

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
