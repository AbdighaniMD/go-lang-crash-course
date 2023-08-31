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

var albumSlices = []album{
	{ID: "1", Title: "Drama setter", Artist: "Eminem", Price: 10.99},
	{ID: "2", Title: "Candy shop", Artist: "50cent", Price: 10.99},
	{ID: "3", Title: "Kill Bill", Artist: "SZA", Price: 10.99},
	{ID: "4", Title: "Princess Diana", Artist: "Ice Spice", Price: 10.99},
	{ID: "5", Title: "Calm Down", Artist: "Rema", Price: 10.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumSlices)
}

func postAlbums(c *gin.Context) {
	var newAlblum album

	err := c.BindJSON(&newAlblum)
	if err != nil {
		return
	}

	albumSlices = append(albumSlices, newAlblum)
	c.IndentedJSON(http.StatusCreated, newAlblum)
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albumSlices {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"Message": "Album not Found",
	})

}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)

	router.Run("localhost:8080")
}
