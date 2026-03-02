package repository

import (
	"backend-api-belajar/model"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type UserRepository interface{
	FindAll()[] model.User
	Save(users model.User)
}

type userRepo struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository  {
	return &userRepo{db: db}
}
func (r *userRepo) FindAll()[]model.User {
	rows, err := r.db.Query("SELECT id,name, age, address FROM users")

	if err!=nil {
		log.Println("error query", err)
		return nil
	}
	defer rows.Close()

	var  users []model.User
	for rows.Next(){
		var u model.User
		rows.Scan(&u.ID, &u.Name,&u.Age,&u.Address)
		if err != nil {
			log.Println("Error scan:", err)
			continue
		}
		users =append(users, u)
	}
	return users
}
func (r *userRepo)Save(user model.User)  {
	r.db.Exec("INSERT INTO users(id,name, age, address) VALUES ($1,$2,$3,$4)", user.ID,user.Name,user.Age,user.Address)
}