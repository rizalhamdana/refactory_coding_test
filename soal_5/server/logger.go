package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

const (
	LOG_FORMAT = "[%s] Success: POST http://%s %s"
)

func LogToFile(dateTime time.Time, ip string, data RequestCounter) {
	f, err := os.OpenFile("server/server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	timeFormattedString := dateTime.Format(time.RFC3339)
	defer f.Close()
	log.SetFlags(0)
	log.SetOutput(f)
	json, _ := json.Marshal(data)
	dataString := string(json)
	log.Printf(LOG_FORMAT, timeFormattedString, ip, dataString)
}
