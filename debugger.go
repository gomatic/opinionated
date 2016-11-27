package main

import (
	_ "expvar"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//
func Debugger() {
	go func() {
		DEBUGGER_PORT := os.Getenv("DEBUGGER_PORT")
		if DEBUGGER_PORT == "" {
			DEBUGGER_PORT = "8008"
		}

		srv := &http.Server{
			Addr:         ":" + DEBUGGER_PORT,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		fmt.Println("listening on " + DEBUGGER_PORT)
		log.Println(srv.ListenAndServe())
	}()
}
