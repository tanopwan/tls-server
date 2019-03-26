package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	srv := &http.Server{
		Addr: ":8888",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{\"resultCode\":\"0\",\"mobile\":\"0803363515\"}"))
		}),
		TLSConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}

	go func() { log.Fatal(srv.ListenAndServeTLS("certificate.pem", "key.pem")) }()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
