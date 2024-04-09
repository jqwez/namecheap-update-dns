package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpBinResponse struct {
	Origin string `json:"origin"`
}

func GetMyIp() (string, error) {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		log.Println("Failed to get IPADDR from httpbin.org/ip")
		log.Println("Falling back to .env")
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	var httpBinResponse HttpBinResponse
	if err := json.NewDecoder(resp.Body).Decode(&httpBinResponse); err != nil {
		log.Println("Failed to decode JSON from httpbin.org/ip")
		return "", err
	}
	return httpBinResponse.Origin, nil
}
