package main

import(
	
	"log"
	"net/http"
	"backend-api-belajar/handler"
	"backend-api-belajar/repository"
	"backend-api-belajar/service"
	"backend-api-belajar/config"
	
)

func main()  {
	db := config.ConnectDB()
	defer db.Close()



	usersRepo := repository.NewUserRepository(db)
	usersService:= service.NewUserService(usersRepo)
	usersHandler:= handler.NewUserHandler(usersService)

	productsRepo := repository.NewProductRepository(db)
	productService:= service.NewProductService(productsRepo)
	productsHandler:= handler.NewProductHandler(productService)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			usersHandler.GetUsers(w,r)
		case http.MethodPost:
			usersHandler.CreateUser(w,r)
		case http.MethodPut:
			usersHandler.UpdateUser(w,r)
		case http.MethodDelete:
			usersHandler.DeleteUser(w,r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
	
		}
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			productsHandler.GetProduct(w,r)
		case http.MethodPost:
			productsHandler.CreateProduct(w,r)
		case http.MethodPut:
			productsHandler.UpdateProduct(w,r)
		case http.MethodDelete:
			productsHandler.DeleteProduct(w,r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
	
		}
	})
	log.Println("server jalan di localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}