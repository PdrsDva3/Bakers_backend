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

// @Summary Delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/delete/{id} [delete]
func (handler HandlerUser) Delete(g *gin.Context) {
	userID := g.Query("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = handler.service.Delete(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"delete": id})
}

// @Summary Change username
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body entities.UserChangeName true "change name"
// @Success 200 {object} int "Success"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/change/name [put]
func (handler HandlerUser) ChangeName(g *gin.Context) {
	var user entities.UserChangeName
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := handler.service.ChangeName(ctx, user.ID, user.Name)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"name": user.Name})
}

// @Summary Change password
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body entities.UserChangePassword true "change name"
// @Success 200 {object} int "Success"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/change/password [put]
func (handler HandlerUser) ChangePassword(g *gin.Context) {
	var user entities.UserChangePassword
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := handler.service.ChangePassword(ctx, user.ID, user.Password)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"change": "success"})
}
