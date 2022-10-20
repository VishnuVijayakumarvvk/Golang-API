package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Transaction struct {
	BatchID         string `json:"Batch ID"`
	EventID         string `json:"Event ID"`
	UserID          string `json:"User ID"`
	Transactioninfo string `json:"Transaction info"`
	Status          bool   `json:"Status"`
}

type Status struct {
	Statuses bool `json:"Status"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/create-transaction", createEmployee)
	http.ListenAndServe(":6100", r)
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API Batch status</h1>"))
}

func createEmployee(w http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tranc Transaction
	var stat Status
	if request.PostFormValue("Batch ID") == "" || request.PostFormValue("Event ID") == "" || request.PostFormValue("Transaction info") == "" || request.PostFormValue("User ID") == "" {
		stat.Statuses = false
		w.Header().Set("context.Type", "application/json")
		json.NewEncoder(w).Encode(stat)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		
		tranc.BatchID = request.PostFormValue("Batch ID")
		tranc.EventID = request.PostFormValue("Event ID")
		tranc.UserID = request.PostFormValue("User ID")
		tranc.Transactioninfo = request.PostFormValue("Transaction info")
		tranc.Status = true
		w.Header().Set("context.Type", "application/json")
		json.NewEncoder(w).Encode(tranc)
	}
	w.WriteHeader(200)
	return

}
