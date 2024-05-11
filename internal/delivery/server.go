package delivery

import (
	"Bakers_backend/docs"
	"Bakers_backend/internal/delivery/handlers"
	"Bakers_backend/internal/repository/admin"
	adminserv "Bakers_backend/internal/service/admin"
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

	//routers.InitRouting(r, db, logger, middlewareStruct, jwtUtils, session, tracer)

	publicRouter := r.Group("/admin")

	userRepo := admin.NewAdminRepo(db)

	publicService := adminserv.InitAdminService(userRepo)
	adminHandler := handlers.InitPublicHandler(publicService)

	publicRouter.POST("/create", adminHandler.CreateAdmin)

	//publicRouter.POST("/login", publicHandler.LoginUser)
	//
	//publicRouter.POST("/refresh", publicHandler.Refresh)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
