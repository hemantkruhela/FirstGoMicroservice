package app

import (
	"log"
	"net/http"

)

func Start() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomres)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
