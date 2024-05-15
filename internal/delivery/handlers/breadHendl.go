package handlers

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
// @Tags bread
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := p.service.BreadCreate(ctx, breadCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get bread
// @Tags bread
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

// @Summary ChangeCount bread
// @Tags bread
// @Accept  json
// @Produce  json
// @Param data body entities.BreadChange true "bread change count (add or sub)"
// @Success 200 {object} int "Successfully change count"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /bread/change [put]
func (p BreadHandler) ChangeCount(c *gin.Context) {
	var change entities.BreadChange

	if err := c.ShouldBindJSON(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	cnt, err := p.service.ChangeBread(ctx, change.BreadID, change.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": change.BreadID, "NewCount": cnt})
}

// @Summary Delete bread
// @Tags bread
// @Accept  json
// @Produce  json
// @Param id query int true "BreadID"
// @Success 200 {object} int "Successfully delete bread"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /bread/delete/{id} [delete]
func (p BreadHandler) DeleteBread(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = p.service.DeleteBread(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": id})
}
