package handlers

import (
	"Bakers_backend/internal/entities"
	"Bakers_backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerUser struct {
	service service.UserServ
}

func InitUserHandler(service service.UserServ) HandlerUser {
	return HandlerUser{
		service: service,
	}
}

// @Summary Create user
// @Tags public
// @Accept  json
// @Produce  json
// @Param data body entities.UserCreate true "user create"
// @Success 200 {object} int "Successfully created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/create [post]
func (handler HandlerUser) CreateUser(g *gin.Context) {
	var newUser entities.UserCreate

	if err := g.ShouldBindJSON(&newUser); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := g.Request.Context()

	id, err := handler.service.Create(ctx, newUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}
