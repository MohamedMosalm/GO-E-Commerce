package category

import (
	"fmt"
	"net/http"

	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"github.com/MohamedMosalm/GO-E-Commerce/utils"
)

type Handler struct {
	Repo CategoryRepository
}

func NewHandler(repository CategoryRepository) *Handler {
	return &Handler{Repo: repository}
}

func (h *Handler) HandleGetcategories(w http.ResponseWriter, r *http.Request) {
	categorys, err := h.Repo.GetAllCategories()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to retrieve categorys"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, categorys)
}

func (h *Handler) HandleGetCategory(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	category, err := h.Repo.GetCategoryByID(uint(categoryID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, category)
}

func (h *Handler) HandleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updates := new(models.Category)
	if err := utils.ParseJSON(r, updates); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	newcategory, err := h.Repo.UpdateCategory(categoryID, updates)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to update category"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, newcategory)
}

func (h *Handler) HandleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.Repo.DeleteCategory(categoryID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	type CategoryPayload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ProductsIds []int  `json:"products_ids"`
	}

	var payload CategoryPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	category := models.Category{
		Name:        payload.Name,
		Description: payload.Description,
	}

	err := h.Repo.CreateCategory(&category, payload.ProductsIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, category)
}
