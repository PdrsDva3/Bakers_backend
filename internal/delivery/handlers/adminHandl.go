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

type AdminHandler struct {
	service service.AdminService
}

func InitAdminHandler(service service.AdminService) AdminHandler {
	return AdminHandler{
		service: service,
	}
}

// @Summary Create admin
// @Tags admin
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := p.service.AdminCreate(ctx, adminCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Login admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param data body entities.AdminLogin true "admin login"
// @Success 200 {object} int "Successfully login user, returning JWT and Session"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/login [post]
func (p AdminHandler) LoginAdmin(c *gin.Context) {
	var adminLogin entities.AdminLogin

	if err := c.ShouldBindJSON(&adminLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := p.service.Login(ctx, adminLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id query int true "AdminID"
// @Success 200 {object} int "Successfully get user, returning JWT and Session"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/{id} [get]
func (p AdminHandler) GetAdmin(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	adm, err := p.service.GetMe(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"admin": adm})
}

// @Summary ChangePWD admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param data body entities.AdminChangePWD true "admin change pwd"
// @Success 200 {object} int "Successfully change pwd, returning JWT and Session"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/change [put]
func (p AdminHandler) ChangePWD(c *gin.Context) {
	var changePWD entities.AdminChangePWD

	if err := c.ShouldBindJSON(&changePWD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err := p.service.ChangePassword(ctx, changePWD.AdminID, changePWD.NewPWD)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"change": "access"})
}

// @Summary Delete admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id query int true "AdminID"
// @Success 200 {object} int "Successfully delete user, returning JWT and Session"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/delete/{id} [delete]
func (p AdminHandler) DeleteAdmin(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = p.service.Delete(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete": id})
}
