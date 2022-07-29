package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Movie struct{
	ID        string `json:"id"`     
	Title     string `json:"title"`     
	Year      int `json:"year"`
	Director *Director string `json:"director"`

}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies = []Movie{
	{ID:"1", Title:"Mokalik", Year:2019, Director: &Directotor{Firstname:"Kunle", Lastname:"Afolayan"} },
	{ID:"2", Title:"King Of Boys",Year:2018, Director: &Director{Firstname:"Kemi", Lastname:"Adetiba"}},
	{ID:"3", Title:"Sugar Rush",Year:2019, Director: &Director{Firstname:"Jade", Lastname:"Osiberu"}},
	{ID:"4", Title:"The Ghost And House OF Truth",Year:2019, Director: &Director{Firstname:"Akin", Lastname:"Omotosho"}},
	{ID:"5", Title:"La Femme Anjola",Year:2020, Director: &Director{Firstname:"Mildred", Lastname:"Okwo"}},
	{ID:"6", Title:"Omo Gheto The Saga",Year:2020, Director: &Director{Firstname:"Funke", Lastname:"Akindele"}},
	{ID:"7", Title:"Aleros Symphony",Year:2011, Director: &Director{Firstname:"Izu", Lastname:"Ojukwu"}},
	{ID:"8", Title:"The Herbert Macaulay Affair",Year:2019, Director: &Director{Firstname:"Imoh", Lastname:"Umoren"}},
	{ID:"9", Title:"The Set Up",Year:2019, Director: &Director{Firstname:"Niyi", Lastname:"Akinmolayan"}},
	{ID:"10", Title:"The Royal Hibiscus Hotel",Year:2017, Director: &Director{Firstname:"Ishaya", Lastname:"Bako"}},
	

}

func getMovies(c *gin.Context){
	c.IndentedJSON(http.StatusOK, movies)
}

func getMovieByID(c *gin.Context){
	ID := c.Param("id")

	for _, item = range movies{
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
		c.IndentedJSON(http.StatusNotFound, gin.H("message": "ID not found"))
}

func postMovie(c *gin.Context){
	var newMovie Movie

	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	movies = append(movies, newMovie)
	
	c.IndentedJSON(http.StatusCreated, newMovie )
}

func deleteMovie(){
	id := c.Param("id")

	for index, item = range movies{
		if item.id == id{
			movies = append(movies[:index], movies[index+1]...)
			break
		}
	}
	c.IndentedJSON(htttp.StatusNotFound, gin.H{"msg": "Movie NOt Found"} )

}

func updateMovie(){
	id := c.Param("id")

	var updateData Movie

	if err := c.BindJSON(&updateData); err != nil {
		return
	}

	for i, item = range movies{
		if item.ID == id{
			movies[i] = updateData
			c.IndentedJSON(http.StatusOK, movies[i])
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "movie not found"})
}

func main() {
	router := gin.Default()

	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovieByID)
	router.POST("/movies", postMovie)
	router.DELETE("/movies/:id", deleteMovie)
	router.PUT("/movies/:id", updateMovie)


	router.Run("localhost:8080")

}
