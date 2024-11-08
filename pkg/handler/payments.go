package handler

import (
	"gogogo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Payment
// @Security ApiKeyAuth
// @Tags payments
// @Description Create Payment
// @ID create-payment
// @Accept json
// @Produce json
// @Param input body gogogo.CreatePaymentInput true "Data"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/payments [post]
func (h *Handler) createPayment(c *gin.Context) {
	var input gogogo.CreatePaymentInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Payments.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllPaymentsResponse struct {
	Data []gogogo.Payment `json:"data"`
}

// @Summary Get All Payments
// @Security ApiKeyAuth
// @Tags payments
// @Description Get All Payments
// @ID get-all-payments
// @Accept json
// @Produce json
// @Success 200 {object} getAllPaymentsResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
func (h *Handler) getAllPayments(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	payments, err := h.services.Payments.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPaymentsResponse{
		Data: payments,
	})
}

// @Summary Get Payment By Id
// @Security ApiKeyAuth
// @Tags payments
// @Description Get Payment By Id
// @ID get-payment-by-id
// @Accept json
// @Produce json
// @Param id path int  true  "Payment ID"
// @Success 200 {object} gogogo.Payment
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/payments/{id} [get]
func (h *Handler) getPaymentById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	payment, err := h.services.Payments.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, payment)
}

// @Summary Get Payment By Order Id
// @Security ApiKeyAuth
// @Tags payments
// @Description Get Payment By Order Id
// @ID get-payment-by-order-id
// @Accept json
// @Produce json
// @Param order_id path int  true  "Order ID"
// @Success 200 {object} gogogo.Payment
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/payments/order/{order_id} [get]
func (h *Handler) getPaymentByOrderId(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	order_id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid order_id param")
		return
	}

	payment, err := h.services.Payments.GetByOrderId(order_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, payment)
}

// @Summary Update Payment
// @Security ApiKeyAuth
// @Tags payments
// @Description Update Payment
// @ID update-payment
// @Accept json
// @Produce json
// @Param id path int  true  "Payment ID"
// @Param input body gogogo.UpdatePaymentInput true "Data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/payments/{id} [put]
func (h *Handler) updatePayment(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input gogogo.UpdatePaymentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Payments.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Payment
// @Security ApiKeyAuth
// @Tags payments
// @Description Delete Payment
// @ID delete-payment
// @Accept json
// @Produce json
// @Param id path int  true  "Payment ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/payments/{id} [delete]
func (h *Handler) deletePayment(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Payments.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
