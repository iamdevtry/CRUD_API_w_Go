package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie = []Movie{
	{
		ID:       "1",
		Isbn:     "12345",
		Title:    "Movie One",
		Director: &Director{Firstname: "John", Lastname: "Doe"},
	},
	{
		ID:       "2",
		Isbn:     "12346",
		Title:    "Movie Two",
		Director: &Director{Firstname: "Steve", Lastname: "Smith"},
	},
}

func getMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, movies)
	}
}

func getMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		for _, movie := range movies {
			if movie.ID == id {
				c.JSON(http.StatusOK, movie)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
	}
}

func addMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		movie.ID = strconv.Itoa(len(movies) + 1)
		movies = append(movies, movie)
		c.JSON(http.StatusOK, movie)
	}
}

func updateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		for index, movie := range movies {
			if movie.ID == id {
				var newMovie Movie
				if err := c.ShouldBindJSON(&newMovie); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				newMovie.ID = id
				movies[index] = newMovie
				c.JSON(http.StatusOK, newMovie)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
	}
}

func deleteMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		for index, movie := range movies {
			if movie.ID == id {
				movies = append(movies[:index], movies[index+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
	}
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/movies", getMovies())
	r.GET("/movies/:id", getMovie())
	r.POST("/movies", addMovie())
	r.PUT("/movies/:id", updateMovie())
	r.DELETE("/movies/:id", deleteMovie())

	r.Run()
}
