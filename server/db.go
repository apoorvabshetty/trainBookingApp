package server

import (
    //"fmt"
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "strconv"
    "trainbooking/models"
)

var DB *gorm.DB

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
}

func InitDB(dbConfig DatabaseConfig) {
	//dsn := "experiment:experiment@tcp(localhost:3306)/recordings?charset=utf8&parseTime=True&loc=Local" 
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbConfig.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
    //dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    /**
	Migrate the schema
	*/
    Migrate()
}

func Migrate() {
    DB.AutoMigrate(&models.Admin{})
    DB.AutoMigrate(&models.Passenger{})
    DB.AutoMigrate(&models.Train{})
}
