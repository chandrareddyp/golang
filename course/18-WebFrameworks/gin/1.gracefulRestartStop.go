package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

/*
example from https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
*/
func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered from panic", r)
		}
	}()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	// gin.ListenAndServe(":8080", router)
	//router.Run(":8081")
	//http.ListenAndServe(":8080", router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("recovered from panic", r)
			}
		}()
		if err := server.ListenAndServe(); err != nil {
			//log.Fatalf("error:%s",err)
			//log.Panic(err) // this logs and calls panic, which will looke for defer statements with recover if any otherwise exits
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	//signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Signal received:", <-quit)
	log.Println("Shutdown Server ...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Microsecond)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
