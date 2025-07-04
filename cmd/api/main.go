package main

import (
	"database/sql"
	"log"
	_ "rest-api-go-gin/docs"
	"rest-api-go-gin/internal/database"
	"rest-api-go-gin/internal/env"

	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"
)

// @title Go Gin Rest API
// @version 1.24.3
// @description A rest API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-123456"),
		models:    models,
	}
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
