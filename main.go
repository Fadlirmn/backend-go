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



	repo := repository.NewUserRepository(db)
	service:= service.NewUserService(repo)
	handler:= handler.NewUserHandler(service)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.GetUsers(w,r)
		case http.MethodPost:
			handler.CreateUser(w,r)
		case http.MethodPut:
			handler.UpdateUser(w,r)
		case http.MethodDelete:
			handler.DeleteUser(w,r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
	
		}
	})
	log.Println("server jalan di localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}