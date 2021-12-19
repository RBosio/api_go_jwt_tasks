package main

import (
	"log"
	"net/http"
	"packages/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	enableCORS(mux)

	mux.HandleFunc("/api/task", handlers.GetTasks).Methods("GET")
	mux.HandleFunc("/api/task/{id:[0-9]}", handlers.GetTask).Methods("GET")
	mux.HandleFunc("/api/task", handlers.NewTask).Methods("POST")
	mux.HandleFunc("/api/task/{id:[0-9]}", handlers.UpdateTask).Methods("PUT")
	mux.HandleFunc("/api/task/{id:[0-9]}", handlers.DeleteTask).Methods("DELETE")

	mux.HandleFunc("/api/login", handlers.Login).Methods("POST")
	mux.HandleFunc("/api/signup", handlers.NewUser).Methods("POST")

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			next.ServeHTTP(w, req)
		})
}
