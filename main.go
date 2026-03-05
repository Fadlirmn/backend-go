package main

import(
	"backend-api-belajar/handler"
	"backend-api-belajar/repository"
	"backend-api-belajar/service"
	"backend-api-belajar/config"

	"github.com/gin-gonic/gin"
	
)

func main()  {
	db := config.ConnectDB()
	r:=gin.Default()


	//users
	usersRepo := repository.NewUserRepository(db)
	usersService:= service.NewUserService(usersRepo)
	usersHandler:= handler.NewUserHandler(usersService)

	//products
	productsRepo := repository.NewProductRepository(db)
	productService:= service.NewProductService(productsRepo)
	productsHandler:= handler.NewProductHandler(productService)

	//routing gin
	userRoutes:= r.Group("/users")
	{
		userRoutes.GET("",usersHandler.GetUsers)
		userRoutes.POST("",usersHandler.CreateUser)
		userRoutes.PUT("",usersHandler.UpdateUser)
		userRoutes.DELETE("",usersHandler.DeleteUser)
	}

	productRoutes:= r.Group("/products")
	{
		productRoutes.GET("",productsHandler.GetProduct)
		productRoutes.POST("",productsHandler.CreateProduct)
		productRoutes.PUT("",productsHandler.UpdateProduct)
		productRoutes.DELETE("",productsHandler.DeleteProduct)
	}
	r.Run(":8080")
}