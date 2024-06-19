package main 

import(
	"net/http"
	"github.com/gin-gonic/gin"
	
)
// blueprint for unmarshalling  json data
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	album{ID: "1", Title: "title1", Artist: "artist1", Price: 10.99},
    album{ID: "2", Title: "title2", Artist: "artist2", Price: 20.99},
    album{ID: "3", Title: "title3", Artist: "artist3", Price: 30.99},
    album{ID: "4", Title: "title4", Artist: "artist4", Price: 40.99},
    album{ID: "5", Title: "title5", Artist: "artist5", Price: 50.99},
}
func main() {
	
	router := gin.Default()
	router.GET("/albums",getAlbums)
	router.POST("/albums",postAlbums)
	router.Run(":8080")
	
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK,albums)
}

func postAlbums(c *gin.Context){
	 var newAlbum album 
	//  call bindJSON to bind the recieved JSON to newalbum.
	if err := c.BindJSON(&newAlbum); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	// Add the new album to the slice
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}