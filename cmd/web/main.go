package main

import (
	"flag"
	"net/http"
	"time"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/loggers"
	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/templating"
)

func main() {

	//Setting the flags of the program, i can control the port and the database information, I'm using
	//MariaDB as a database software right now.
	addr := flag.String("addr", ":9090", "http network address: A number that defines the addr of our application")
	dsn := flag.String("dsn", "web:passwd1234@/snippetbox?parseTime=true", "MySQL dsn")
	debug := flag.Bool("debug", false, "Set to true to activate the debug mode")
	flag.Parse()

	cache, err := templating.NewTemplateCache()
	if err != nil {
		loggers.ErrorLog().Fatal(err)
	}

	// This just fetches the configurations at compile time. It is useful for getting the flags and cache running and ready,
	app := config.ExportConfig(*dsn, *debug, cache)

	app.Infolog.Printf("Starting application on port %v \n", *addr) //Logging on start

	//err := http.ListenAndServe(addr, Router())
	//if err != nil {
	//	log.Fatal("error on http.ListenAndServe function, check the addr and mux variables.")
	//}
	srv := &http.Server{
		Handler:  Router(app),
		Addr:     *addr,
		ErrorLog: app.Errorlog,

		WriteTimeout: 5 * time.Second,  //Prevents the handler data from taking too long to write.
		ReadTimeout:  10 * time.Second, //This one prevents slow client attacks
		IdleTimeout:  1 * time.Minute,  //This one is purely done for the server performance (i/e closing conections) when the user logs out due to unexpected events.

		TLSConfig: app.TLSconfig,
	}
	app.Errorlog.Fatal(srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"))
}
