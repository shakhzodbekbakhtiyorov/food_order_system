package handler

import (
	"fmt"
	"gogogo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary CreateCategory
// @Security ApiKeyAuth
// @Tags categories
// @Description Create Category
// @ID create-category
// @Accept json
// @Produce json
// @Param input body gogogo.CreateCategoryInput true "Data"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/categories [post]
func (h *Handler) createCategory(c *gin.Context) {
	var input gogogo.CreateCategoryInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Categories.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllCategoriesResponse struct {
	Data []gogogo.Category `json:"data"`
}

// @Summary GetAllCategories
// @Security ApiKeyAuth
// @Tags categories
// @Description Get All Category
// @ID get-all-categories
// @Accept json
// @Produce json
// @Success 200 {object} getAllCategoriesResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/categories [get]
func (h *Handler) getAllCategories(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	categories, err := h.services.Categories.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCategoriesResponse{
		Data: categories,
	})
}

// @Summary GetCategoryById
// @Security ApiKeyAuth
// @Tags categories
// @Description Get Category By Id
// @ID get-category-by-id
// @Accept json
// @Produce json
// @Param id path int  true  "Category ID"
// @Success 200 {object} gogogo.Category
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/categories/{id} [get]
func (h *Handler) getCategoryById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	fmt.Println(c.Param("id"))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	category, err := h.services.Categories.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, category)
}

// @Summary UpdateCategory
// @Security ApiKeyAuth
// @Tags categories
// @Description Update Category
// @ID update-category
// @Accept json
// @Produce json
// @Param id path int  true  "Category ID"
// @Param input body gogogo.CreateCategoryInput true "CreateCategoryInput"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/categories/{id} [put]
func (h *Handler) updateCategory(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input gogogo.CreateCategoryInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Categories.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary GetCategoryById
// @Security ApiKeyAuth
// @Tags categories
// @Description Delete Category
// @ID delete-category
// @Accept json
// @Produce json
// @Param id path int  true  "Category ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/categories/{id} [delete]
func (h *Handler) deleteCategory(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Categories.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
