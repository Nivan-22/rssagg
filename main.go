package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nivan-22/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Started!!")

	godotenv.Load(".env")
	
	portS := os.Getenv("PORT")
	if portS == "" {
		log.Fatal("PORT is not found")
	}
	dataB := os.Getenv("DB_URL")
	if dataB == "" {
		log.Fatal("Database is not found")
	}
	conn, err := sql.Open("postgres", dataB)
	if err != nil {
		log.Fatal("Error in opening database")
	}
   
	apiCfg := apiConfig{
		DB:  database.New(conn),
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

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	router.Mount("/v1", v1Router)
    v1Router.Get("/users",apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portS,
	}
	log.Printf("Server starts on :%v", portS)

	log.Fatal(srv.ListenAndServe())

	fmt.Println("Port:", portS)
}
