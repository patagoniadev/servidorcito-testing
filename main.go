package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Temperatura de la pava
type Temperatura struct {
	Temperatura string `json:"temperatura"`
}

func getTemperatura(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "70")
}

func calentar(w http.ResponseWriter, r *http.Request) {
	var temperatura Temperatura

	if err := json.NewDecoder(r.Body).Decode(&temperatura); err != nil {
		fmt.Fprintf(w, "Formato Incorrecto")
		return
	}

	fmt.Fprintf(w, "Calentando... %s\n", temperatura.Temperatura)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/temperatura-actual", getTemperatura).Methods("POST")
	router.HandleFunc("/calentar", calentar).Methods("POST")

	log.Fatal(http.ListenAndServe(":6080", router))
}
