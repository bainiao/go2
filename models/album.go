package models

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAllAlbums() []Album {
	return Albums
}

func TestFunc() string {
	return "This is a test function in models package"
}

func PostAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.Bind(&newAlbum); err != nil {
		return
	}
	Albums = append(Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, Albums)
}

func filterAlbumsByID(id string) ([]Album, bool) {
	var res []Album
	for _, a := range Albums {
		if a.ID == id {
			res = append(res, a)
		}
	}
	if len(res) > 0 {
		return res, true
	}
	return []Album{}, false
}
func filterAlbumsByTitle(title string) (Album, bool) {
	for _, a := range Albums {
		if strings.Contains(strings.ToLower(a.Title), strings.ToLower(title)) {
			return a, true
		}
	}
	return Album{}, false
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, ok := filterAlbumsByID(id)
	if ok {
		c.IndentedJSON(http.StatusOK, album)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func GetAlbumByTitle(c *gin.Context) {
	title := c.Query("title")
	albums, ok := filterAlbumsByTitle(title)
	if ok {
		c.IndentedJSON(http.StatusOK, albums)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}
