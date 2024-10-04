package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "golang.org/x/crypto/bcrypt"
    "trainbooking/models"
    "trainbooking/server"
)

// AdminLogin handles admin login
func AdminLogin(c *gin.Context) {
    var admin models.Admin
    if err := c.ShouldBindJSON(&admin); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var storedAdmin models.Admin
    if err := server.DB.Where("username = ?", admin.Username).First(&storedAdmin).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        // return
    }

    if err := server.DB.Where("password = ?", admin.Password).First(&storedAdmin).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // if err := bcrypt.CompareHashAndPassword([]byte(storedAdmin.Password), []byte(admin.Password)); err != nil {
    //     c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
    //     return
    // }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// CreateAdmin creates a new admin user
func CreateAdmin(c *gin.Context) {
    var admin models.Admin
    if err := c.ShouldBindJSON(&admin); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
    // admin.Password = string(hashedPassword)

    if err := server.DB.Create(&admin).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
        return
    }
    c.JSON(http.StatusCreated, admin)
}
