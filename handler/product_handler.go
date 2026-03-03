package handler

import (
	"backend-api-belajar/model"
	"backend-api-belajar/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductHandler struct{
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService)*ProductHandler  {
	return &ProductHandler{service: s}
}

func (h *ProductHandler)GetProduct(w http.ResponseWriter, r *http.Request)  {
	products:= h.service.GetProduct()
	json.NewEncoder(w).Encode(products)
}
func (h *ProductHandler)CreateProduct(w http.ResponseWriter, r *http.Request)  {
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	h.service.CreateProduct(product)
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler)UpdateProduct(w http.ResponseWriter, r *http.Request)  {
	StringId:= r.URL.Query().Get("id")
	id, err:= strconv.Atoi(StringId)
	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is Required"))
		return
	
	}
	var product model.Product
	// 2. Decode JSON (Hanya simpan hasil error-nya ke 'err')
	// JANGAN masukkan hasil Decode ke dalam 'id'
	err = json.NewDecoder(r.Body).Decode(&product) 
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON body"))
		return
	}

	// 3. Panggil service dengan 'id' yang sudah jadi int
	err = h.service.UpdateProduct(id, product)

		w.Write([]byte("Product updated successfully"))
}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    StringId := r.URL.Query().Get("id")
    id, err := strconv.Atoi(StringId)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("ID is Required and must be a number"))
        return // Cukup return saja, jangan 'return nil' atau 'return err'
    }

    err = h.service.DeleteProduct(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Failed to delete product"))
        return 
    }

    w.Write([]byte("Product Deleted Successfully"))
}