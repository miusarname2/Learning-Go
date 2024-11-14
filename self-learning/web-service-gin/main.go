package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Album 1", Artist: "Artist 1", Price: 10.99},
	{ID: "2", Title: "Album 2", Artist: "Artist 2", Price: 15.99},
	{ID: "3", Title: "Album 3", Artist: "Artist 1", Price: 8.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080

}
