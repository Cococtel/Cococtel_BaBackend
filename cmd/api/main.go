package main

import (
	"database/sql"
	"fmt"
	"github.com/Cococtel/Cococtel_BaBackend/internal/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	godotenv.Load(".env")
	db, err := createDB()
	if err != nil {
		panic(err)
	}
	eng := gin.Default()
	router := http.InitRouter(eng, db)
	router.MapRoutes()
	if err := eng.Run(); err != nil {
		panic(err)
	}
}
func createDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&tls=true", dbUser, dbPassword, dbHost, dbName)

	return sql.Open("mysql", connectionString)
}
