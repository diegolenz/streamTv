package main

import (
	"log"
	"net/http"
	"stream-api/pkg/api"
	"stream-api/pkg/daos"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Configura os cabeçalhos CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite todas as origens
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Trata as requisições OPTIONS (pré-voo do CORS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	_, err := daos.Connect()

	if err != nil {
		panic(err)
	}
	log.Println(daos.DbConection)

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permite todas as origens
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Aplica o middleware

	//router.HandleFunc("/", home.ServeHTTP)
	router.HandleFunc("/upload", api.UploadVideoHandler).Methods("POST")
	router.HandleFunc("/find-video/{index}", api.GetVideoHandler).Methods("GET")
	c.Handler(router)
	log.Println("API started.")
	log.Fatal(http.ListenAndServe(":8000", router))

}
