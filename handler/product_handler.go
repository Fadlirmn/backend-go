package handler

import (
	"backend-api-belajar/model"
	"backend-api-belajar/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService)*ProductHandler  {
	return &ProductHandler{service: s}
}

func (h *ProductHandler)GetProduct(c *gin.Context)  {
	products:= h.service.GetProduct()
	c.JSON(http.StatusOK, products)
}
func (h *ProductHandler)CreateProduct(c *gin.Context)  {
	var product model.Product
	json.NewDecoder(c.Request.Body).Decode(&product)

	h.service.CreateProduct(product)
	c.JSON(http.StatusAccepted,product)
}

func (h *ProductHandler)UpdateProduct(c *gin.Context)  {
	StringId:= c.Query("id")
	id, err:= strconv.Atoi(StringId)
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"ID MUST NUMBER"})
		return
	
	}
	var product model.Product
	// 2. Decode JSON (Hanya simpan hasil error-nya ke 'err')
	// JANGAN masukkan hasil Decode ke dalam 'id'
	err = json.NewDecoder(c.Request.Body).Decode(&product) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Json Body"})
		return
	}

	// 3. Panggil service dengan 'id' yang sudah jadi int
	err = h.service.UpdateProduct(id, product)

		c.JSON(http.StatusOK, err)
}
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
    StringId := c.Query("id")
    id, err := strconv.Atoi(StringId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":"ID is Required and must be a number"})
        return // Cukup return saja, jangan 'return nil' atau 'return err'
    }

    err = h.service.DeleteProduct(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error":"ID is Required and must be a number"})
        return 
    }

     c.JSON(http.StatusOK, gin.H{"success":"Done Deleted Products"})
}