package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Details struct {
	Name     string `json:"name"`
	NodeName string `json:"nodename"`
}

func displayDetails(res http.ResponseWriter, req *http.Request) {
	var details = map[string]*Details{
		"01": &Details{Name: "fralix", NodeName: "node01"},
	}
	res.Header().Set("Content-Type", "application/json")
	outgoingJSON, error := json.Marshal(details)
	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(outgoingJSON))
}

func HelloService() {
	msgrouter := mux.NewRouter()
	msgrouter.HandleFunc("/details", displayDetails).Methods("GET")
	http.ListenAndServe(":9090", msgrouter)
}
