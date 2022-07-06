package handler

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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

	id, err := h.services.Purchase.AddPurchase(purchaseData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
