package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go_crud/pkg/container"
	"go_crud/pkg/member"
	"go_crud/pkg/mysql"
	"go_crud/pkg/tweet"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var shutdownSignal = make(chan os.Signal)
	signal.Notify(shutdownSignal, syscall.SIGTERM)
	signal.Notify(shutdownSignal, syscall.SIGINT)
	ctx := context.Background()

	app := &container.Container{
		Debug:  true,
		Router: mux.NewRouter(),
	}

	mysql.Configure(app)
	tweet.LoadRoutes(app)
	member.LoadRoutes(app)

	go func() {
		if err := http.ListenAndServe(":8010", app.Router); err != nil {
			log.Println(fmt.Errorf("Error serving the app: %w", err))
		}
	}()

	fmt.Println("Server listen on localhost:8010")
	<-shutdownSignal
	_, cancel := context.WithCancel(ctx)
	defer cancel()
}
