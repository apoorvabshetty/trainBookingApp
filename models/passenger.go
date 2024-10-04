package models

type Passenger struct {
    ID      uint   `json:"id" gorm:"primaryKey"`
    Name    string `json:"name"`
    Email   string `json:"email" gorm:"unique"`
    Age     int    `json:"age"`
    Gender  string `json:"gender"`
}
