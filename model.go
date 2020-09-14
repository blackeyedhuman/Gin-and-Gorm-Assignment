
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Model struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Length string `json:"length"`
}

var allmodels []Model

func GetAllModels(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allmodels)

}
func GetModel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	from_browser := mux.Vars(r)
	for _, values_in_allmodels := range allmodels {
		if values_in_allmodels.Id == from_browser["id"] {
			json.NewEncoder(w).Encode(values_in_allmodels)
			return
		}

	}

}
func CreateModel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new_model Model
	json.NewDecoder(r.Body).Decode(&new_model)
	new_model.Id = strconv.Itoa(len(allmodels) + 1)
	allmodels = append(allmodels, new_model)
	json.NewEncoder(w).Encode(allmodels)

}
func DeleteModel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	from_browser := mux.Vars(r)
	for i, values_in_allmodels := range allmodels {
		if values_in_allmodels.Id == from_browser["id"] {
			allmodels = append(allmodels[:i], allmodels[i+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(allmodels)

}
func UpdateModel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	from_browser := mux.Vars(r)
	for i, values_in_allmodels := range allmodels {
		if values_in_allmodels.Id == from_browser["id"] {
			allmodels = append(allmodels[:i], allmodels[i+1:]...)
			var new_updated_model Model
			json.NewDecoder(r.Body).Decode(&new_updated_model)
			new_updated_model.Id = from_browser["id"]
			allmodels = append(allmodels, new_updated_model)
			json.NewEncoder(w).Encode(allmodels)
			return
		}

	}
	json.NewEncoder(w).Encode(allmodels)
}

func main() {
	router := mux.NewRouter()
	allmodels = append(allmodels, Model{Id: "1", Name: "Learn Go quick", Length: "2 hours"})
	router.HandleFunc("/models", GetAllModels).Methods("GET")
	router.HandleFunc("/models/{id}", GetModel).Methods("GET")
	router.HandleFunc("/model", CreateModel).Methods("POST")
	router.HandleFunc("/models/{id}", DeleteModel).Methods("DELETE")
	router.HandleFunc("/models/{id}", UpdateModel).Methods("POST")

	log.Fatal(http.ListenAndServe(":2000", router))
}