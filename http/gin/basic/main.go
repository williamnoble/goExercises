package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	ID     string  `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Artist string  `json:"artist,omitempty"`
	Price  float64 `json:"price,omitempty"`
}

var Albums []Album

func main() {
	Albums = []Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Kind of Blue", Artist: "Miles Davis", Price: 20.99},
		{ID: "3", Title: "Enter the Wu-Tang", Artist: "Wu-Tang Clan", Price: 17.99},
	}

	router := gin.Default()
	router.GET("/Albums", getAlbums)
	router.GET("/Albums/:id", getAlbumsById)
	router.POST("/Albums", postAlbums)
	_ = router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}

func postAlbums(c *gin.Context) {
	var album Album
	if err := c.BindJSON(album); err != nil {
		return
	}

	Albums = append(Albums, album)
}

func getAlbumsById(c *gin.Context) {
	id := c.Params.ByName("id")
	for _, album := range Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
