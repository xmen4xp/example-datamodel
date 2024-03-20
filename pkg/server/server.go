package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var (
	server *http.Server
	router *mux.Router
)

func StartServer(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done() // let main know we are done cleaning up

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
}

func CreateServer() {
	router = mux.NewRouter()
	server = &http.Server{Addr: ":8080", Handler: router}
}

func StopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("shutdown of server failed with error %v", err)
	}
}

func AddRoute(route string, f func(http.ResponseWriter, *http.Request), methods ...string) {
	router.HandleFunc(route, f).Methods(methods...)
}
