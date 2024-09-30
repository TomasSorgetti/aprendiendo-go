package models


type User struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"size:100"`
    Email string `gorm:"size:100"`
    Password string `gorm:"size:100"`
}
