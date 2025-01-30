package main

import (
	"database/sql"
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
	dns := os.Getenv("DB_DNS")
	return sql.Open("mysql", dns)
}
