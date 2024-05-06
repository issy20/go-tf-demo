package main

import (
	"encoding/json"
	"go-tf-demo/cmd/cmd"
	"go-tf-demo/config"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	var cfg config.Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(cfg.DBName)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/health", healthHandler)

	server.ListenAndServe()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	response := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Stage   string `json:"stage"`
	}{
		Code:    200,
		Message: "OK",
		Stage:   cfg.Stage,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
