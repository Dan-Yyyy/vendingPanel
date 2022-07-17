package handler

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/message"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// @Summary AddPurchase
// @Tags Purchase
// @Description Добавление единицы закупки
// @ID addPurchase
// @Accept json
// @Produce json
// @Param input body models.Purchase true "Данные о закупке"
// @Success 200 {object} int "Индентификатор закупки"
// @Failure 400,404 {object} responseError
// @Failure 500 {object} responseError
// @Failure default {object} responseError
// @Router /api/purchase [post]
func (h *Handler) addPurchase(c *gin.Context) {
	var purchaseData models.Purchase

	err := c.BindJSON(&purchaseData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	purchaseData.CreatedAt = time.Now()
	purchaseData.UpdatedAt = time.Now()
	purchaseData.UserId = userId

	purchaseId, err := h.services.Purchase.AddPurchase(purchaseData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	stockData := models.Stock{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		PurchaseId: purchaseId,
	}

	_, err = h.services.Stock.AddStock(stockData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": purchaseId,
	})
}

// @Summary GetPurchase
// @Tags Purchase
// @Description Получение единицы закупки
// @ID getPurchase
// @Accept json
// @Produce json
// @Success 200 {object} models.Purchase "Объект закупки"
// @Failure 400,404 {object} responseError
// @Failure 500 {object} responseError
// @Failure default {object} responseError
// @Router /api/purchase/:id [get]
func (h *Handler) getPurchase(c *gin.Context) {
	purchaseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), message.InvalidUserId)
		return
	}

	purchase, err := h.services.Purchase.GetPurchase(purchaseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	c.JSON(http.StatusOK, purchase)
}

// @Summary GetPurchases
// @Tags Purchase
// @Description Получение единицы закупки
// @ID getPurchases
// @Accept json
// @Produce json
// @Success 200 {object} []models.Purchase "Массив закупок"
// @Failure 400,404 {object} responseError
// @Failure 500 {object} responseError
// @Failure default {object} responseError
// @Router /api/purchase [get]
func (h *Handler) getPurchases(c *gin.Context) {
	purchases, err := h.services.Purchase.GetPurchases()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	c.JSON(http.StatusOK, purchases)
}

// @Summary UpdatePurchases
// @Tags Purchase
// @Description Обновление единицы закупки
// @ID getPurchases
// @Accept json
// @Produce json
// @Param input body models.Purchase true "Обновленые данные закупки"
// @Success 200 {object} responseStatus
// @Failure 400,404 {object} responseError
// @Failure 500 {object} responseError
// @Failure default {object} responseError
// @Router /api/purchase/:id [post]
func (h *Handler) updatePurchase(c *gin.Context) {
	var purchaseData models.Purchase

	err := c.BindJSON(&purchaseData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	purchaseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), message.InvalidUserId)
		return
	}

	purchaseData.UpdatedAt = time.Now()

	err = h.services.Purchase.UpdatePurchases(purchaseId, purchaseData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	c.JSON(http.StatusOK, responseStatus{"ok"})
}

// @Summary DeletePurchase
// @Tags Purchase
// @Description Удаление единицы закупки
// @ID getPurchase
// @Accept json
// @Produce json
// @Success 200 {object} responseStatus
// @Failure 400,404 {object} responseError
// @Failure 500 {object} responseError
// @Failure default {object} responseError
// @Router /api/purchase/:id [delete]
func (h *Handler) deletePurchase(c *gin.Context) {
	purchaseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), message.InvalidUserId)
		return
	}

	err = h.services.Purchase.DeletePurchase(purchaseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	c.JSON(http.StatusOK, responseStatus{"ok"})
}
