package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"github.com/MohamedMosalm/GO-E-Commerce/services/auth"
	"github.com/MohamedMosalm/GO-E-Commerce/utils"
	"github.com/gorilla/mux"
)

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

		if _, err := h.Repo.GetUserByEmail(payload.Email); err == nil {
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

func (h *Handler) HandleUpdatePassword(w http.ResponseWriter, r *http.Request) {

	type UpdatePasswordPayload struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}

	var payload UpdatePasswordPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.Repo.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user is not found, error: %v", err))
		return
	}
	if auth.ComparePasswords(user.HashedPassword, []byte(payload.Password)) {
		newHashedPassword, err := auth.HashPassword(payload.NewPassword)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		user.HashedPassword = newHashedPassword

		if err := h.Repo.SaveIntoDB(user); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to update password: %v", err))
		}

		utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "password updated successfully"})
	} else {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("incorrect password"))
	}
}

func (h *Handler) HandleForgotPassword(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Email string `json:"email"`
	}
	var payload Request
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.Repo.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("User not found"))
	}

	token, err := utils.GenerateToken()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to generate token"))
		return
	}

	resetToken := &models.PasswordResetToken{
		Email:     payload.Email,
		Token:     token,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	if err := h.PasswordResetTokenRepo.CreatePasswordResetToken(resetToken); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to save token : %v", err))
		return
	}

	resetURL := utils.GenerateResetURL(r, token)
	fmt.Println(resetURL)

	if err := utils.SendEmail([]string{payload.Email}, resetURL); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to send email"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Password reset email sent"})
}

func (h *Handler) HandleResetPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resetToken := vars["resetToken"]

	token, err := h.PasswordResetTokenRepo.FindResetToken(resetToken)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Token is invalid or expired. Please try again"))
		return
	}

	user, err := h.Repo.GetUserByEmail(token.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Token is invalid or expired. Please try again"))
		return
	}

	type Request struct {
		Password string `json:"password"`
	}

	var payload Request
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Error hashing password: %v", err))
		return
	}
	user.HashedPassword = hashedPassword
	err = h.Repo.SaveIntoDB(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Error hashing password: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "password reset successfully"})
}
