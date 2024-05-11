package delivery

import (
	"Bakers_backend/docs"
	"Bakers_backend/internal/delivery/handlers"
	"Bakers_backend/internal/repository/admin"
	"Bakers_backend/internal/repository/bread"
	adminserv "Bakers_backend/internal/service/admin"
	breadserv "Bakers_backend/internal/service/bread"
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

	adminRouter := r.Group("/admin")

	adminRepo := admin.InitAdminRepo(db)
	adminService := adminserv.InitAdminService(adminRepo)
	adminHandler := handlers.InitAdminHandler(adminService)

	adminRouter.POST("/create", adminHandler.CreateAdmin)
	adminRouter.POST("/login", adminHandler.LoginAdmin)
	adminRouter.GET("/:id", adminHandler.GetAdmin)
	adminRouter.DELETE("/delete/:id", adminHandler.DeleteAdmin)
	adminRouter.PUT("/change", adminHandler.ChangePWD)

	breadRouter := r.Group("/bread")

	breadRepo := bread.InitBreadRepo(db)
	breadService := breadserv.InitBreadService(breadRepo)
	breadHandler := handlers.InitBreadHandler(breadService)

	breadRouter.POST("/create", breadHandler.CreateBread)
	breadRouter.GET("/:id", breadHandler.GetBread)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
