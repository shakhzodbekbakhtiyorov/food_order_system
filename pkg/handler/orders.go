package handler

import (
	"gogogo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Order
// @Security ApiKeyAuth
// @Tags orders
// @Description Create Order
// @ID create-order
// @Accept json
// @Produce json
// @Param input body gogogo.CreateOrderInput true "Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/orders [post]
func (h *Handler) createOrder(c *gin.Context) {
	var input gogogo.CreateOrderInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Orders.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllOrdersResponse struct {
	Data []gogogo.Order `json:"data"`
}

// @Summary Get All Orders
// @Security ApiKeyAuth
// @Tags orders
// @Description Get All Orders
// @ID get-all-orders
// @Accept json
// @Produce json
// @Success 200 {object} getAllOrdersResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/orders [get]
func (h *Handler) getAllOrders(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	orders, err := h.services.Orders.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

// @Summary Get Order By Id
// @Security ApiKeyAuth
// @Tags orders
// @Description Get Order By Id
// @ID get-order-by-id
// @Accept json
// @Produce json
// @Param id path int  true  "Order ID"
// @Success 200 {object} gogogo.Order
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/orders/{id} [get]
func (h *Handler) getOrderById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	product, err := h.services.Orders.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Update Order
// @Security ApiKeyAuth
// @Tags orders
// @Description Update Order
// @ID update-order
// @Accept json
// @Produce json
// @Param id path int  true  "Order ID"
// @Param input body gogogo.UpdateOrderInput true "Data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/orders/{id} [put]
func (h *Handler) updateOrder(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input gogogo.UpdateOrderInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Orders.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Order
// @Security ApiKeyAuth
// @Tags orders
// @Description Delete Order
// @ID delete-order
// @Accept json
// @Produce json
// @Param id path int  true  "Order ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/orders/{id} [delete]
func (h *Handler) deleteOder(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Orders.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
