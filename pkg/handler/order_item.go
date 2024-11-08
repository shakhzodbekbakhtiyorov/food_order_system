package handler

import (
	"gogogo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Order Item
// @Security ApiKeyAuth
// @Tags orderitems
// @Description Create Order Item
// @ID createorderitem
// @Accept json
// @Produce json
// @Param input body gogogo.CreateOrderItem true "Data"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order_items [post]
func (h *Handler) createOrderItem(c *gin.Context) {
	var input gogogo.CreateOrderItem

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.OrderItems.CreateOrderItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Update Order Item
// @Security ApiKeyAuth
// @Tags orderitems
// @Description Update Order Item
// @ID update-orderitem
// @Accept json
// @Produce json
// @Param id path int  true  "Order Item ID"
// @Param input body gogogo.UpdateOrderItem true "Data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order_items/{id} [put]
func (h *Handler) updateOrderItem(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input gogogo.UpdateOrderItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.OrderItems.EditOrderItem(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Order Item
// @Security ApiKeyAuth
// @Tags orderitems
// @Description Delete Order Item
// @ID delete-orderitem
// @Accept json
// @Produce json
// @Param id path int  true  "Order Item ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/order_items/{id} [delete]
func (h *Handler) deleteOderItem(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.OrderItems.DeleteOrderItem(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
