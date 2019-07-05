package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alicebob/miniredis"
)

func main() {
	server := miniredis.NewMiniRedis()
	err := server.StartAddr("0.0.0.0:6379")
	if err != nil {
		panic(err)
	}
	log.Printf("miniredis started on 127.0.0.1:6379")
	mainloop()
	defer server.Close()
}

func mainloop() {
    exitSignal := make(chan os.Signal)
    signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
    <-exitSignal

	log.Printf("exiting...")
}
