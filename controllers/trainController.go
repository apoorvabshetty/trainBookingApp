package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "trainbooking/models"
    "trainbooking/server"
)

func CreateTrain(c *gin.Context) {
    var train models.Train
    if err := c.ShouldBindJSON(&train); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := server.DB.Create(&train).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create train"})
        return
    }
    c.JSON(http.StatusCreated, train)
}

func ListTrains(c *gin.Context) {
    var trains []models.Train
    if err := server.DB.Find(&trains).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trains"})
        return
    }
    c.JSON(http.StatusOK, trains)
}

func EditTrain(c *gin.Context) {
    var train models.Train
    id := c.Param("id")

    if err := server.DB.First(&train, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Train not found"})
        return
    }

    if err := c.ShouldBindJSON(&train); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := server.DB.Save(&train).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update train"})
        return
    }
    c.JSON(http.StatusOK, train)
}

func DeleteTrain(c *gin.Context) {
    id := c.Param("id")
    if err := server.DB.Delete(&models.Train{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete train"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

func ViewTrainAvailability(c *gin.Context) {
    var trains []models.Train
    if err := server.DB.Find(&trains).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch train availability"})
        return
    }
    c.JSON(http.StatusOK, trains)
}
