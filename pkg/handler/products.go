package handler

import (
	"gogogo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Product
// @Security ApiKeyAuth
// @Tags products
// @Description Create Product
// @ID create-product
// @Accept json
// @Produce json
// @Param input body gogogo.CreateProductInput true "Data"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/products [post]
func (h *Handler) createProduct(c *gin.Context) {
	var input gogogo.CreateProductInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Products.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllProductsResponse struct {
	Data []gogogo.Product `json:"data"`
}

// @Summary Get All Products
// @Security ApiKeyAuth
// @Tags products
// @Description Get All Products
// @ID get-all-products
// @Accept json
// @Produce json
// @Success 200 {object} getAllProductsResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/products [get]
func (h *Handler) getAllProducts(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	products, err := h.services.Products.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllProductsResponse{
		Data: products,
	})
}

// @Summary Get Product By Id
// @Security ApiKeyAuth
// @Tags products
// @Description Get Product By Id
// @ID get-product-by-id
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {object} gogogo.Product
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/products/{id} [get]
func (h *Handler) getProductById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	product, err := h.services.Products.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Update Product
// @Security ApiKeyAuth
// @Tags products
// @Description Update Product
// @ID update-product
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Param input body gogogo.UpdateProductInput true "Data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/products/{id} [put]
func (h *Handler) updateProduct(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input gogogo.UpdateProductInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Products.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Product
// @Security ApiKeyAuth
// @Tags products
// @Description Delete Product
// @ID delete-product
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/products/{id} [delete]
func (h *Handler) deleteProduct(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Products.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
