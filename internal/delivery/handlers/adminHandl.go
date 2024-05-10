package handlers

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminHandler struct {
	service service.AdminService
}

func InitPublicHandler(service service.AdminService) AdminHandler {
	return AdminHandler{
		service: service,
	}
}

// @Summary Create admin
// @Tags public
// @Accept  json
// @Produce  json
// @Param data body entities.AdminCreate true "admin create"
// @Success 200 {object} int "Successfully created user, returning JWT and Session"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/create [post]
func (p AdminHandler) CreateAdmin(c *gin.Context) {
	var adminCreate entities.AdminCreate

	if err := c.ShouldBindJSON(&adminCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	id, err := p.service.AdminCreate(ctx, adminCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
