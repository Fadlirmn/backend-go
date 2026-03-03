package repository

import (
	"backend-api-belajar/model"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"math/rand"
	"time"
	"github.com/oklog/ulid/v2"
)

type UserRepository interface{
	FindAll()[] model.User
	Save(users model.User)
	Update(id string, user model.User)error
	Delete(id string)error
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
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())),0)
	id := ulid.MustNew(ulid.Timestamp(t),entropy).String()


	_ , err:= r.db.Exec("INSERT INTO users(id, name, age, address) VALUES ($1,$2,$3,$4)", id, user.Name,user.Age,user.Address)
	if err != nil{
		log.Println("GAGAL nambah database: ", err)
	}
}

func (r *userRepo) Update(id string, user model.User)error {
	_, err:= r.db.Exec("UPDATE users SET name = $1, age = $2, address = $3 WHERE id= $4", user.Name, user.Age, user.Address, id)
	return err
}

func (r *userRepo)Delete(id string) error {
	_, err:= r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}