package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Message struct {
	message string `json:"message"`
}

type RequestCounter struct {
	Counter int    `json:"counter"`
	Header  string `json:"header"`
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthCheck)

	fmt.Println("SERVER LISTEN AT PORT: 8080")

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := RequestCounter{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Header = r.Header.Get("X-RANDOM")

		LogToFile(time.Now(), GetLocalIP(), data)
		w.WriteHeader(http.StatusCreated)

	} else {
		w.Header().Set("Content-Type", "application/json")
		message := Message{message: "Bad Request"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
	}

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{message: "Server is Up"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
