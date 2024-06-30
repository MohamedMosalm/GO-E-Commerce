package api

import (
	"net/http"

	"github.com/MohamedMosalm/GO-E-Commerce/services/auth"
	"github.com/MohamedMosalm/GO-E-Commerce/services/category"
	"github.com/MohamedMosalm/GO-E-Commerce/services/product"
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

func RegisterProductsRoutes(r *APIServer, apiRouter *mux.Router) {
	productRepo := product.NewGromProductRepository(r.db)
	productHandler := product.NewHandler(productRepo)
	apiRouter.HandleFunc("/products", productHandler.HandleGetProducts).Methods(http.MethodGet)
	apiRouter.HandleFunc("/products", productHandler.HandleCreateProduct).Methods(http.MethodPost)
	apiRouter.HandleFunc("/products/{ID}", productHandler.HandleGetProduct).Methods(http.MethodGet)
	apiRouter.HandleFunc("/products/{ID}", productHandler.HandleDeleteProduct).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/products/{ID}", productHandler.HandleUpdateProduct).Methods(http.MethodPatch)
}

func RegisterCategoriesRoutes(r *APIServer, apiRouter *mux.Router) {
	categoryRepo := category.NewGromCategoryRepository(r.db)
	categoryHandler := category.NewHandler(categoryRepo)
	apiRouter.HandleFunc("/categories", categoryHandler.HandleGetcategories).Methods(http.MethodGet)
	apiRouter.HandleFunc("/categories", categoryHandler.HandleCreateCategory).Methods(http.MethodPost)
	apiRouter.HandleFunc("/categories/{ID}", categoryHandler.HandleGetCategory).Methods(http.MethodGet)
	apiRouter.HandleFunc("/categories/{ID}", categoryHandler.HandleDeleteCategory).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/categories/{ID}", categoryHandler.HandleUpdateCategory).Methods(http.MethodPatch)
}
