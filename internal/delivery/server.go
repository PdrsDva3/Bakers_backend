package delivery

import (
	"Bakers_backend/docs"
	"Bakers_backend/internal/delivery/handlers"
	"Bakers_backend/internal/repository/admin"
	"Bakers_backend/internal/repository/bread"
	"Bakers_backend/internal/repository/user"
	adminserv "Bakers_backend/internal/service/admin"
	breadserv "Bakers_backend/internal/service/bread"
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
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// todo: имплементировать cors middleware
	//middlewareStruct := middleware.InitMiddleware(logger, jwtUtils, session)
	//r.Use(middlewareStruct.CORSMiddleware())

	userRouter := r.Group("/user")

	userRepo := user.InitUserRepository(db)
	userService := userserv.InitUserService(userRepo)
	userHandler := handlers.InitUserHandler(userService)

	userRouter.POST("/create", userHandler.CreateUser)
	userRouter.POST("/login", userHandler.Login)
	userRouter.GET("/:id", userHandler.Get)

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
	breadRouter.DELETE("/delete/:id", breadHandler.DeleteBread)
	breadRouter.PUT("/change", breadHandler.ChangeCount)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
