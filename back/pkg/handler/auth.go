package handler

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var userData models.User

	if err := c.BindJSON(&userData); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorisation.CreateUser(userData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
