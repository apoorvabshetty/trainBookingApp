package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "trainbooking/models"
    "trainbooking/server"
)

func CreatePassenger(c *gin.Context) {
    var passenger models.Passenger
    if err := c.ShouldBindJSON(&passenger); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := server.DB.Create(&passenger).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create passenger"})
        return
    }
    c.JSON(http.StatusCreated, passenger)
}

func ListPassengers(c *gin.Context) {
    var passengers []models.Passenger
    if err := server.DB.Find(&passengers).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch passengers"})
        return
    }
    c.JSON(http.StatusOK, passengers)
}

func EditPassenger(c *gin.Context) {
    var passenger models.Passenger
    id := c.Param("id")

    if err := server.DB.First(&passenger, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found"})
        return
    }

    if err := c.ShouldBindJSON(&passenger); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := server.DB.Save(&passenger).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update passenger"})
        return
    }
    c.JSON(http.StatusOK, passenger)
}

func DeletePassenger(c *gin.Context) {
    id := c.Param("id")
    if err := server.DB.Delete(&models.Passenger{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete passenger"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}
