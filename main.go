package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HrMinSec struct {
	hour   string
	minute string
	second string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello world")
}
