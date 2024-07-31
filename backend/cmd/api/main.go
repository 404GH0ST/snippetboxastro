package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/404GH0ST/snippetboxastro/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	model *models.SnippetModel
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbDsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPass,
		dbName,
	)

	db, err := openDB(dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &Application{
		model: &models.SnippetModel{DB: db},
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	e.GET("/api/view/latest", app.snippetLatest)
	e.GET("/api/view/:id", app.snippetView)
	e.POST("/api/create", app.snippetCreate)

	e.Logger.Error(e.Start(host))
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(
		context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
