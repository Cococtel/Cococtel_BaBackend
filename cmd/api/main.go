package main

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/Cococtel/Cococtel_BaBackend/internal/http"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

func main() {
	godotenv.Load(".env")
	db, err := createDB()
	if err != nil {
		panic(err)
	}
	err = db.Ping()
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
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_connector.go: %s environment variable not set.", k)
		}
		return v
	}
	env := mustGetenv("ENV")
	if env != "dev" {
		instanceConnectionName := mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
		usePrivate := os.Getenv("PRIVATE_IP")
		d, err := cloudsqlconn.NewDialer(context.Background())
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
		}
		var opts []cloudsqlconn.DialOption
		if usePrivate != "" {
			opts = append(opts, cloudsqlconn.WithPrivateIP())
		}
		mysql.RegisterDialContext("cloudsqlconn",
			func(ctx context.Context, addr string) (net.Conn, error) {
				return d.Dial(ctx, instanceConnectionName, opts...)
			})
	}
	dns := os.Getenv("DB_DNS")
	return sql.Open("mysql", dns)
}
