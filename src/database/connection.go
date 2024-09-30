package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error al cargar el archivo .env:", err)
    }

    var dbUser = os.Getenv("DB_USER")
    var dbPassword = os.Getenv("DB_PASSWORD")
    var dbName = os.Getenv("DB_NAME")
    var dbHost = os.Getenv("DB_HOST")
    var dbPort = os.Getenv("DB_PORT")

    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

	log.Println("Conexión exitosa a la base de datos MySQL")
}