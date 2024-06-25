package user

import (
	"fmt"
	"net/http"

	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"github.com/MohamedMosalm/GO-E-Commerce/utils"
)

type Handler struct {
	Repo                   UserRepository
	PasswordResetTokenRepo PasswordResetTokenRepository
}

func NewHandler(repository UserRepository, passwordResetTokenRepo PasswordResetTokenRepository) *Handler {
	return &Handler{Repo: repository, PasswordResetTokenRepo: passwordResetTokenRepo}
}

func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to retrieve users: %v", err))
		return
	}
	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user, err := h.Repo.GetUserByID(uint(userID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updates := new(models.User)
	if err := utils.ParseJSON(r, updates); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	newuser, err := h.Repo.UpdateUser(userID, updates)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to update user"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, newuser)
}

func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.Repo.DeleteUser(userID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusNoContent, nil)
}
