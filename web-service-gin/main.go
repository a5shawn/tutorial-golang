package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func getAlbums(context *gin.Context) {
	context.JSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	err := context.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.JSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(context *gin.Context) {
	id := context.Param("id")
	for _, album := range albums {
		if album.ID == id {
			context.JSON(http.StatusOK, album)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
