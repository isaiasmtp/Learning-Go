package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/isaiasmtp/Learning-Go/core/beer"
	"github.com/isaiasmtp/Learning-Go/web/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data/beer.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service := beer.NewService(db)
	r := mux.NewRouter()
	//middlewares - código que vai ser executado em todas as requests
	//aqui podemos colocar logs, inclusão e validação de cabeçalhos, etc
	n := negroni.New(
		negroni.NewLogger(),
	)
	//handlers
	handlers.MakeBeerHandlers(r, n, service)

	//static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	r.PathPrefix("/static/").Handler(n.With(
		negroni.Wrap(http.StripPrefix("/static/", fileServer)),
	)).Methods("GET", "OPTIONS")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// used to health check, will return 200
	})

	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + os.Getenv("HTTP_PORT"),
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
