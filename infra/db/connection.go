package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
func GetConnectionString()string{
	return "user=postgres password=2314 host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConncection()(*sqlx.DB,error){
	dbsource := GetConnectionString()

	dbcon, err := sqlx.Connect("postgres",dbsource)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}
	return dbcon,nil
}