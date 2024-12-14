package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerInfo struct {
	Hostname string `json:"hostname"`
	ServerId string `json:"server_id"`
	StartTime time.Time `json:"start_time"`
	RequestTime time.Time `json:"requst_time"`
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverId := os.Getenv("SERVER_ID") 
	
	if serverId == "" {
		serverId = "tidak diketahui"
	}

	startTime := time.Now()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		// object
		info := ServerInfo{
			Hostname: hostname,
			ServerId: serverId,
			StartTime: startTime,
			RequestTime: time.Now(),
			Message: "Request berhasil diproses",
		}

		// header json
		w.Header().Set("Content-Type", "application/json")

		// json
		json.NewEncoder(w).Encode(info)
	})

	log.Printf("Server %s berjalan di port %s", serverId, port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}