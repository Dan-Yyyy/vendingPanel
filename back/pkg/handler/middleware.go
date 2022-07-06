package handler

import (
	"errors"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/message"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, message.EmptyAuthHeader, message.EmptyAuthHeader)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, message.InvalidAuthHeader, message.InvalidAuthHeader)
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, message.TokenIsEmpty, message.TokenIsEmpty)
		return
	}

	userId, err := h.services.Authorisation.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error(), message.TokenExpired)
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New(message.UserIdNotFound)
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New(message.InvalidTypeOfUserId)
	}

	return idInt, nil
}
