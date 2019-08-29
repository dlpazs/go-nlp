package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"restapi/controller"
	"github.com/joho/godotenv"
)


func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main(){
	fmt.Println("Starting application")	
	
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).
		Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).
		Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).
		Methods("GET")
	r.HandleFunc("/summarize", controller.SummarizationHandler).
		Methods("POST")
	
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	
	log.Fatal(http.ListenAndServe(":8000", r))
}