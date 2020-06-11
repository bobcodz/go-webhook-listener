package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/listen", Listen).Methods("POST")

	fmt.Println("Listening on port 8000")

	http.ListenAndServe(":8000", router)
}

func Listen(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

