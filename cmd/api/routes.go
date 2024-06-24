package api

import (
	"net/http"

	"github.com/MohamedMosalm/GO-E-Commerce/services/auth"
	"github.com/MohamedMosalm/GO-E-Commerce/services/user"
	"github.com/gorilla/mux"
)

func RegisterUsersRoutes(r *APIServer, apiRouter *mux.Router) {
	userRepo := user.NewGromUserRepository(r.db)
	userHandler := user.NewHandler(userRepo)
	apiRouter.HandleFunc("/register", userHandler.HandleRegister).Methods(http.MethodPost)
	apiRouter.HandleFunc("/login", userHandler.HandleLogin).Methods(http.MethodPost)
	apiRouter.Handle("/logout", auth.AuthMiddleware(userHandler.HandleLogout)).Methods(http.MethodGet)
}
