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

type HandlerUser struct {
	service service.UserServ
}

func InitUserHandler(service service.UserServ) HandlerUser {
	return HandlerUser{
		service: service,
	}
}

// @Summary Create user
// @Tags user
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := handler.service.Create(ctx, newUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully get user"
// @Failure 400 {object} map[string]string "Invalid UserID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/{id} [get]
func (handler HandlerUser) Get(g *gin.Context) {
	//todo доделать до middleware
	userID := g.Query("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := g.Request.Context()
	user, err := handler.service.Get(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary Login user
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body entities.UserLogin true "user login"
// @Success 200 {object} int "Successfully login user"
// @Failure 400 {object} map[string]string "Invalid password"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/login [post]
func (handler HandlerUser) Login(g *gin.Context) {
	var User entities.UserLogin

	if err := g.ShouldBindJSON(&User); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := handler.service.Login(ctx, User)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}
