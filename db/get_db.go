package db

import (
	"github.com/go-pg/pg/v10"
)

func GetDb(dbConnectionString string, dbUser string, dbPassword string, dbName string) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     dbConnectionString,
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})
}
