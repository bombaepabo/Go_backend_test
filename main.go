package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")
	portstring := os.Getenv("PORT")
	if(portstring == ""){
		log.Fatal("Port is not found in the environment")
	}
	router := chi.NewRouter()
	server := &http.Server{
		Handler: router,
		Addr: ":" + portstring,
	}
	log.Printf("Server starting on port %v ",portstring)
	err :=server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Port :",portstring)
}