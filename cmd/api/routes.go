package api

import (
	"net/http"

	"github.com/MohamedMosalm/GO-E-Commerce/services/auth"
	"github.com/MohamedMosalm/GO-E-Commerce/services/user"
	"github.com/gorilla/mux"
)

func RegisterUsersRoutes(r *APIServer, apiRouter *mux.Router) {
	userRepo := user.NewGromUserRepository(r.db)
	passwordResetTokenRepo := user.NewGromPasswordResetTokenRepository(r.db)
	userHandler := user.NewHandler(userRepo, passwordResetTokenRepo)
	apiRouter.HandleFunc("/register", userHandler.HandleRegister).Methods(http.MethodPost)
	apiRouter.HandleFunc("/login", userHandler.HandleLogin).Methods(http.MethodPost)
	apiRouter.HandleFunc("/forgotPassword", userHandler.HandleForgotPassword).Methods(http.MethodPost)
	apiRouter.HandleFunc("/resetPassword/{resetToken}", userHandler.HandleResetPassword).Methods(http.MethodPost)
	apiRouter.HandleFunc("/updatePassword", auth.AuthMiddleware(userHandler.HandleUpdatePassword)).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/logout", auth.AuthMiddleware(userHandler.HandleLogout)).Methods(http.MethodGet)

	apiRouter.HandleFunc("/users", userHandler.HandleGetUsers).Methods(http.MethodGet)
	apiRouter.HandleFunc("/users/{ID}", userHandler.HandleGetUser).Methods(http.MethodGet)
	apiRouter.HandleFunc("/users/{ID}", userHandler.HandleDeleteUser).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/users/{ID}", userHandler.HandleUpdateUser).Methods(http.MethodPatch)
}
