package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *sql.DB
}

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, errsql := sql.Open("mysql", "root:@(db:3306)/lab3")
	if errsql != nil {
		log.Fatal(errsql)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		db:       db,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Вызов нового метода app.routes()
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
