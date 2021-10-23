package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	store "github.com/fatihsimsek/go-case-study/internal/store"
)

func main() {
	serverAddress := getServerAddr()

	quit := make(chan bool)
	defer close(quit)

	initPackages(quit)

	log.Printf("Server Starting at: %s\n", serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getServerAddr() string {
	value, exists := os.LookupEnv("PORT")
	if !exists {
		value = "8090"
	}
	return ":" + value
}

func initPackages(quit chan bool) {
	service := store.NewService(store.NewRepository())
	store.Init(service)
	go runSchedulers(service, quit)
}

func runSchedulers(service store.Service, quit chan bool) {
	ticker := time.NewTicker(time.Duration(getScheduleInterval()) * time.Second)
	for {
		select {
		case <-quit:
			ticker.Stop()
			return
		case <-ticker.C:
			log.Println("Ticker ticking...")
			service.Flush()
		}
	}
}

func getScheduleInterval() int {
	value, exists := os.LookupEnv("SCHEDULEINTERVAL")
	if !exists {
		value = "10"
	}
	result, _ := strconv.Atoi(value)
	return result
}
