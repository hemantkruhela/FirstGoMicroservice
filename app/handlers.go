package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Customer ...
type Customer struct {
	Name string `json:"Your_NAME" xml:"Naam"`
	City string `json:"CITY_NAME" xml:"city"`
	Pin  string `json:"PINCODE" xml:"PinCODE"`
}

func getAllCustomres(w http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{"Ashish", "Delhi", "110023"},
		{"Dinesh", "Delhi", "110032"},
	}

	if request.Header.Get("Accept") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func greet(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Greetings and welcome to GOLANG")
}
