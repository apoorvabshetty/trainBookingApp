package models

type Admin struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username" gorm:"unique"`
    Password string `json:"password"`
}
