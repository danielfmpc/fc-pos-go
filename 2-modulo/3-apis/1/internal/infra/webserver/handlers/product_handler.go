package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pos-go-api/internal/dto"
	"pos-go-api/internal/entity"
	"pos-go-api/internal/infra/database"
	pkg "pos-go-api/pkg/entity"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create product godoc
// @Summary 		Create product
// @Description Create a new product
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			request body 		 	dto.CreateProductInputDto true "product request"
// @Success 		201
// @Failure 		400 		{object} 	Error
// @Failure 		500 		{object} 	Error
// @Router 			/products 				[post]
// @Security ApiKeyAuth
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get product by id godoc
// @Summary 		Get product by id
// @Description Get a product by id
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			id 			path 			string true "product id" Format(uuid)
// @Success 		200 		{object} 	entity.Product
// @Failure 		400 		{object} 	Error
// @Failure 		500 		{object} 	Error
// @Router 			/products/{id} 		[get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := pkg.ParseID(id)
	if id == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update product godoc
// @Summary 		Update product
// @Description Update a product
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			id 			path 			string true "product id" Format(uuid)
// @Param 			request body 			dto.CreateProductInputDto true "product request"
// @Success 		200
// @Failure 		400 		{object} 	Error
// @Failure 		500 		{object} 	Error
// @Router 			/products/{id} 		[put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	idParsed, err := pkg.ParseID(id)
	if id == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID = idParsed

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete product godoc
// @Summary 		Delete product
// @Description Delete a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id 			path 			string true "product id" Format(uuid)
// @Success 		200
// @Failure 		400 		{object} 	Error
// @Failure 		500 		{object} 	Error
// @Router 			/products/{id} 		[delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := pkg.ParseID(id)
	if id == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// List products godoc
// @Summary 		List products
// @Description List all products
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			page 			query 		string false "page number"
// @Param 			limit 		query 		string false "limit number"
// @Param 			sort 			query 		string false "sort by"
// @Success 		200 			{object} 	[]entity.Product
// @Failure 		400 			{object} 	Error
// @Failure 		500 			{object} 	Error
// @Router 			/products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if products == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
