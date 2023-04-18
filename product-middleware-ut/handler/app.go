package handler

import (
	"dts-fga-swswg/product-middleware-ut/database"
	"dts-fga-swswg/product-middleware-ut/repository/product_repository/product_pg"
	"dts-fga-swswg/product-middleware-ut/repository/user_repository/user_pg"
	"dts-fga-swswg/product-middleware-ut/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	var port = "8080"
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	productRepo := product_pg.NewProductPG(db)

	productService := service.NewProductService(productRepo)

	productHandler := NewProductHandler(productService)

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, productRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
	}

	productRoute := route.Group("/products")
	{
		productRoute.POST("/", authService.Authentication(), productHandler.CreateProduct)

		productRoute.PUT("/:productId", authService.Authentication(), authService.Authorization(), productHandler.UpdateProductById)
	}

	route.Run(":" + port)
}
