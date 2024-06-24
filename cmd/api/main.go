package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (r *APIServer) Run() error {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(loggingMiddleware)

	RegisterUsersRoutes(r, apiRouter)

	router.Use(mux.CORSMethodMiddleware(router))

	fmt.Println("Listening on", r.addr)
	return http.ListenAndServe(r.addr, router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.RequestURI, time.Since(start))
	})
}
