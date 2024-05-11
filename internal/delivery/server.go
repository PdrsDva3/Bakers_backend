package delivery

import (
	"Bakers_backend/docs"
	"Bakers_backend/internal/delivery/handlers"
	"Bakers_backend/internal/repository/admin"
	"Bakers_backend/internal/repository/user"
	adminserv "Bakers_backend/internal/service/admin"
	userserv "Bakers_backend/internal/service/user"
	"Bakers_backend/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, logger *logger.Logs) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// todo: имплементировать cors middleware
	//middlewareStruct := middleware.InitMiddleware(logger, jwtUtils, session)
	//r.Use(middlewareStruct.CORSMiddleware())

	adminRouter := r.Group("/admin")
	userRouter := r.Group("/user")

	adminRepo := admin.NewAdminRepo(db)
	userRepo := user.InitUserRepository(db)

	adminService := adminserv.InitAdminService(adminRepo)
	userService := userserv.InitUserService(userRepo)

	adminHandler := handlers.InitPublicHandler(adminService)
	userHandler := handlers.InitUserHandler(userService)

	adminRouter.POST("/create", adminHandler.CreateAdmin)

	userRouter.POST("/create", userHandler.CreateUser)
	userRouter.POST("/login", userHandler.Login)
	userRouter.GET("/:id", userHandler.Get)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
