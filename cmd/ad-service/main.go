package main

import (
	"fmt"
	"github.com/bl4ckf1sher/ad-service/internal/application/handlers"
	"github.com/bl4ckf1sher/ad-service/internal/config"
	"github.com/bl4ckf1sher/ad-service/internal/infrastructure/postgres"
	"github.com/bl4ckf1sher/ad-service/internal/infrastructure/repositories"
	"github.com/bl4ckf1sher/ad-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {

	viper.SetConfigName("config.local")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	cfg := &config.DB{}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
		os.Exit(500)
		panic("Error")
	}
	connection := postgres.NewConnect(cfg)
	UserRepo := repositories.NewUserRepository(connection)
	UserService := service.NewUsersService(*UserRepo)
	UserHandler := handlers.NewUsersHandler(*UserService)

	router := gin.Default()
	router.GET("/user", UserHandler.GetUsers)
	router.POST("/user", UserHandler.CreateUser)
	router.GET("/user/:id", UserHandler.GetUserById)
	router.PUT("/user/:id", UserHandler.UpdateUser)
	router.DELETE("/user/:id", UserHandler.DeleteUser)

	router.Run("localhost:8080")
}
