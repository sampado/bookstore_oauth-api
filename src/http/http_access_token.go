package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sampado/bookstore_oauth-api/src/domain/access_token"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid JSON body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
