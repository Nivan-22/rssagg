package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Started!!")
    
	godotenv.Load(".env")
	portS := os.Getenv("PORT")
	if portS == ""{
		log.Fatal("PORT is not found")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router :=  chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	router.Mount("/v1",v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portS,
	}
	log.Printf("Server starts on :%v",portS)
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
    }
    fmt.Println("Port:",portS)
}