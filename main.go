package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HamzahAA15/records_api/db"
	"github.com/HamzahAA15/records_api/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/tanimutomo/sqlfile"
)

const (
	Host     = "postgres"
	Port     = 5432
	User     = "root"
	Password = "password"
	Dbname   = "records_db"
)

func main() {
	app := fiber.New()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)

	dbx, err := db.ConnectPostgres(context.Background(), psqlInfo)
	if err != nil {
		panic(err)
	}
	defer dbx.Close()

	s := sqlfile.New()
	s.Directory("./db/migrations")
	s.Directory("./db/seeds")
	_, err = s.Exec(dbx.DB)
	if err != nil {
		log.Printf("Failed to apply migrations & seeder, run manually from sql scripts in ./db: %v", err)
	}

	internal.SetupRoutes(app, dbx)

	app.Listen(":8080")
}
