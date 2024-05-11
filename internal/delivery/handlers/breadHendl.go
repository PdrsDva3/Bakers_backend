package handlers

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BreadHandler struct {
	service service.BreadService
}

func InitBreadHandler(service service.BreadService) BreadHandler {
	return BreadHandler{
		service: service,
	}
}

// @Summary Create bread
// @Tags public
// @Accept  json
// @Produce  json
// @Param data body entities.BreadBase true "bread create"
// @Success 200 {object} int "Successfully created bread, returning JWT and Session"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /bread/create [post]
func (p BreadHandler) CreateBread(c *gin.Context) {
	var breadCreate entities.BreadBase

	if err := c.ShouldBindJSON(&breadCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	id, err := p.service.BreadCreate(ctx, breadCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get bread
// @Tags public
// @Accept  json
// @Produce  json
// @Param id query int true "BreadID"
// @Success 200 {object} int "Successfully get bread, returning JWT and Session"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /bread/{id} [get]
func (p BreadHandler) GetBread(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	brd, err := p.service.GetBread(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bread": brd})
}
