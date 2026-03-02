package config

import(
	"database/sql"
	"fmt"
	"log"
	"time"

	_"github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB()  *sql.DB{
	dsn:= "postgres://sumbul:admin123@localhost:5432/belajar_go_db"

	db, err := sql.Open("pgx",dsn)
	if err!= nil {
		log.Fatal("Gagal Koneksi ke driver",err)
	}

	err = db.Ping()
	if err!=nil {
		log.Fatal("database tdk respon", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	fmt.Print("✅ berhasil koneksi ke postgresql")
	return db
}