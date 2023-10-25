package main

import (
	"flag"
	"net/http"
	"time"
)

func main() {

	addr := flag.String("addr", ":9090", "http network address: A number that defines the addr of our application")
	flag.Parse()

	inflog := Infolog()
	errlog := ErrorLog()

	app := GetApp()
	app.ErrorLog = errlog
	app.InfoLog = inflog

	app.InfoLog.Printf("Starting application on port %v \n", *addr)

	//err := http.ListenAndServe(addr, Router())
	//if err != nil {
	//	log.Fatal("error on http.ListenAndServe function, check the addr and mux variables.")
	//}
	srv := &http.Server{
		Handler:      Router(),
		Addr:         *addr,
		ErrorLog:     errlog,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	app.ErrorLog.Fatal(srv.ListenAndServe())
}
