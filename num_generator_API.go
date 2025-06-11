package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Answer struct {
	Number int `json:"number"`
}

func main() {
	router := http.NewServeMux()
	NewHandler(router)
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Listening on port 8081")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func NewHandler(r *http.ServeMux) {
	r.HandleFunc("/rand_num", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		answer := Answer{
			Number: rand.Intn(6) + 1,
		}
		if err := json.NewEncoder(w).Encode(answer); err != nil {
			panic(err)
		}
	})
}
