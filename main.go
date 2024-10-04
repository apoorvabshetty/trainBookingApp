package main

import (
    // "github.com/gin-gonic/gin"
    "trainbooking/config"
    "trainbooking/controllers"
    "trainbooking/server"
    "fmt"
)

func main() {
    // Load config
    config.LoadConfig()

    dbConfig := server.DatabaseConfig{
        Host:     config.AppConfig.Database.Host,
        Port:     config.AppConfig.Database.Port,
        User:     config.AppConfig.Database.User,
        Password: config.AppConfig.Database.Password,
        Name:     config.AppConfig.Database.Name,
    }
    server.InitDB(dbConfig)
    router := server.InitRouter()

    router.POST("/admin/login", controllers.AdminLogin)
    router.POST("/admin/create", controllers.CreateAdmin)

    
    router.POST("/admin/passengers", controllers.CreatePassenger)
    router.GET("/admin/passengers", controllers.ListPassengers)
    router.PUT("/admin/passengers/:id", controllers.EditPassenger)
    router.DELETE("/admin/passengers/:id", controllers.DeletePassenger)

    router.POST("/admin/trains", controllers.CreateTrain)
    router.GET("/admin/trains", controllers.ListTrains)
    router.PUT("/admin/trains/:id", controllers.EditTrain)
    router.DELETE("/admin/trains/:id", controllers.DeleteTrain)
    router.GET("/admin/trains/availability", controllers.ViewTrainAvailability)

    router.Run(":" + fmt.Sprint(config.AppConfig.Server.Port))
}



