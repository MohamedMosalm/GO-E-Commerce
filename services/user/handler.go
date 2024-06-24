package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"github.com/MohamedMosalm/GO-E-Commerce/services/auth"
	"github.com/MohamedMosalm/GO-E-Commerce/utils"
)

type Handler struct {
	Repo UserRepository
}

func NewHandler(repository UserRepository) *Handler {
	return &Handler{Repo: repository}
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if !utils.IsValidEmail(payload.Email) || len(payload.Password) < 4 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("please provide a valid credentials"))
		return
	}

	user, err := h.Repo.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	if !auth.ComparePasswords(user.HashedPassword, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("incorrect password"))
		return
	}

	token, err := auth.CreateJWT(payload.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User logged in successfully"})
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-24 * time.Hour),
		HttpOnly: true,
	})

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "logged out successfully"})
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload models.User

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if utils.IsValidEmail(payload.Email) && len(payload.Password) >= 4 && payload.FirstName != "" && payload.LastName != "" {

		if _, err := h.Repo.GetUserByEmail(payload.Email); err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("email already exists"))
			return
		}

		HashedPassword, err := auth.HashPassword(payload.Password)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		err = h.Repo.CreateUser(&models.User{
			FirstName:      payload.FirstName,
			LastName:       payload.LastName,
			Email:          payload.Email,
			HashedPassword: HashedPassword,
		})

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		token, err := auth.CreateJWT(payload.Email)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
	} else {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("please provide a valid credentials"))
		return
	}
}
