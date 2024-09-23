package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
}
var albums = []album{
	{ID: "1", Title: "Blue", Artist: "John"},
	{ID: "2", Title: "Green", Artist: "Sarah"},
	{ID: "3", Title: "Red", Artist: "Bob"},	
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}