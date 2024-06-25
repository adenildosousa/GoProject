package main

import (
	"auth-service/internal/adapter/http"
	"auth-service/internal/core/domain"
	"auth-service/internal/core/service"
	"auth-service/pkg/jwt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	// dsn := "sqlserver://HIAGYU:password@localhost:1433?database=authdb"
	dsn := "Server=HIAGYU;Database=WA_DB;Trusted_Connection=True;Encrypt=False"
	db, err := gorm.Open("mssql", dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()
	db.AutoMigrate(&domain.User{})

	userRepository := sqlserver.NewUserRepository(db)
	userService := service.NewUserService(userRepository, jwt.NewJWTManager("your_secret_key"))

	r := gin.Default()

	http.NewHandler(r, userService)

	r.Run(":8081")

}
