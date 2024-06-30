package product

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"github.com/MohamedMosalm/GO-E-Commerce/utils"
)

type Handler struct {
	Repo ProductRepository
}

func NewHandler(repository ProductRepository) *Handler {
	return &Handler{Repo: repository}
}

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("GetProducts handler called")
	products, err := h.Repo.GetAllProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to retrieve products"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product, err := h.Repo.GetProductByID(uint(productID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *Handler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updates := new(models.Product)
	if err := utils.ParseJSON(r, updates); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	newProduct, err := h.Repo.UpdateProduct(productID, updates)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to update Product"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, newProduct)
}

func (h *Handler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := utils.GetiD(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.Repo.DeleteProduct(productID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	type ProductPayload struct {
		Name          string  `json:"name"`
		Description   string  `json:"description"`
		Price         float64 `json:"price"`
		Stock         int     `json:"stock"`
		CategoriesIds []int   `json:"categories_ids"`
	}

	var payload ProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product := models.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Stock:       payload.Stock,
	}

	err := h.Repo.CreateProduct(&product, payload.CategoriesIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, product)
}
