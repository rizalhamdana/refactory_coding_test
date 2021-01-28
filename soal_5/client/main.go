package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {

	isUp := healthCheck()
	if !isUp {
		fmt.Println("SERVER IS DOWN")
		return
	}
	fmt.Println("SERVER IS UP")
	makeRequest(1)

	for {
		count := 2
		time.Sleep(60 * time.Second)
		go makeRequest(count)
		count++
	}

}

func makeRequest(count int) error {
	client := &http.Client{}
	localIp := "http://" + GetLocalIP() + ":8080"
	requestBody := strings.NewReader(`
		{
			"counter" : ` + fmt.Sprintf("%d", count) + `
		}	
	`)
	header := "daslhquwq"

	req, err := http.NewRequest(http.MethodPost, localIp, requestBody)
	if err != nil {
		return err
	}
	req.Header.Set("X-RANDOM", header)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("FAILED SEND REQUEST")

	}
	fmt.Println(time.Now().String(), "REQUEST SENT TO SERVER")
	return nil
}

func healthCheck() bool {
	localIp := "http://" + GetLocalIP() + ":8080"
	resp, err := http.Get(localIp + "/health")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("SERVER IS DOWN")
		return false
	}
	return true

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
