package db

import (
	"eccomerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
func GetConnectionString(cnf *config.DBcnf) string {
	sslMode := "disable"
	if cnf.EnableSSL {
		sslMode = "require"
	}

	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cnf.DBUser,
		cnf.Password,
		cnf.Host,
		cnf.DBPort,
		cnf.DBname,
		sslMode,
	)
}

func NewConncection(cnf *config.DBcnf)(*sqlx.DB,error){
	dbsource := GetConnectionString(cnf)

	dbcon, err := sqlx.Connect("postgres",dbsource)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}
	return dbcon,nil
}