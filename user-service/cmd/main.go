package main

import (
	"log"
	"user-service/internal/adapter/http"
	"user-service/internal/core/domain"
	"user-service/internal/core/service"
	"user-service/pkg/jwt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	dsn := "sqlserver://username:password@localhost:1433?database=userdb"
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

	r.Run(":8082")
}
