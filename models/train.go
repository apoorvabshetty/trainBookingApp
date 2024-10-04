package models

type Train struct {
    ID             uint   `json:"id" gorm:"primaryKey"`
    TrainNumber    string `json:"train_number" gorm:"unique"`
    TrainName      string `json:"train_name"`
    Route          string `json:"route"`
    Schedule       string `json:"schedule"`
    TotalSeats     int    `json:"total_seats"`
    AvailableSeats  int    `json:"available_seats"`
}
