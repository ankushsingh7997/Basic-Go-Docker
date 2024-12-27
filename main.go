package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Greeting(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello there"))

}
func Testing(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello there"))

}

var port string = ":8080"

func main() {

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	http.HandleFunc("/", Greeting)
	http.HandleFunc("/test", Testing)

	srv := &http.Server{Addr: port,
		Handler:      nil,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		fmt.Println("Server is running on port %s", port)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server gracefully stopped")
}
