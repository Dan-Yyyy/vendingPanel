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

// @Summary SignUp
// @Tags Auth
// @Description Создание нового пользователя
// @ID create-account
// @Accept json
// @Produce json
// @Param input body models.User true "Данные пользователя"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} responseError
// @Failure 500 {object} responseError
// @Failure default {object} responseError
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
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

// @Summary SignIn
// @Tags Auth
// @Description Авторизация пользователя
// @ID login
// @Accept json
// @Produce json
// @Param input body signInData true "Данные пользователя"
// @Success 200 {string} string "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTcwNDYxNTIsImlhdCI6MTY1NzAwMjk1MiwidXNlcl9pZCI6MX0.xhsaDIdJ8BBnp4sh_dBfHjVq8TMFL_AruUDdGxIhBpU"
// @Failure 400,404 {object} responseError
// @Failure 500 {object} responseError
// @Failure default {object} responseError
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var userData signInData

	if err := c.BindJSON(&userData); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorisation.GenerateToken(userData.Email, userData.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": id,
	})
}
