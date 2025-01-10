package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"secondProject/db"
	"secondProject/models"
	"strconv"
	"time"
)

const (
	movieTableName = "movie"
)

var cache = make(map[int64]models.Movie, 500)

func CreateMovie(c *gin.Context) {
	var movieReq models.MovieReq

	if err := c.ShouldBindJSON(&movieReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide valid body request"})
		return
	}

	movie := models.Movie{
		Name:   movieReq.Name,
		Rating: movieReq.Rating,
		Genre:  movieReq.Genre,
	}

	var cnt int64
	db.GetDb().Model(models.Movie{}).Where("name = ?", movieReq.Name).Count(&cnt)
	if cnt > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("movie %s is already exist", movieReq.Name)})
		return
	}

	if err := db.GetDb().Create(&movie).Error; err != nil {
		log.Printf("error to create movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": models.MovieId{
		ID: movie.ID,
	}})
}

func GetMovie(c *gin.Context) {
	var movieFilter models.MovieFilter

	if err := c.Bind(&movieFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide body request"})
	}

	var movie []models.Movie

	DB := db.GetDb().Model(&models.Movie{})

	if movieFilter.Name != "" {
		DB = DB.Where("name LIKE ?", "%"+movieFilter.Name+"%")
	}
	if movieFilter.Rating != nil {
		DB = DB.Where("rating = ?", movieFilter.Rating)
	}

	if movieFilter.Genre != "" {
		DB = DB.Where("genre = ?", movieFilter.Genre)
	}

	if movieFilter.DateFrom != nil {
		t, _ := time.Parse("2006-01-02T15:04", *movieFilter.DateFrom)
		DB = DB.Where("created_at >= ?", t)
	}

	if movieFilter.DateTo != nil {
		t, _ := time.Parse("2006-01-02T15:04", *movieFilter.DateTo)
		DB = DB.Where("created_at <", t)
	}

	page := 1
	limit := 3

	if movieFilter.Page != nil && *movieFilter.Page > 0 {
		page = *movieFilter.Page
	}

	var count int64
	DB.Count(&count)
	DB = DB.Offset((page - 1) * limit).Limit(limit)
	if err := DB.Order("id asc").Find(&movie).Error; err != nil {
		log.Printf("error to get movies : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
	}

	c.JSON(http.StatusOK, gin.H{"movies": movie, "total_count": count, "page": page})
}

func UpdateMovie(c *gin.Context) {
	var movieUpd models.Movie

	if err := c.ShouldBindJSON(&movieUpd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide body request"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "#"})
		return
	}

	cnt := models.Movie{
		Name:   movieUpd.Name,
		Rating: movieUpd.Rating,
		Genre:  movieUpd.Genre,
	}

	if err = db.GetDb().Table(movieTableName).Where("id = ?", id).Updates(cnt).Error; err != nil {
		log.Printf("error to delete movie by id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": cnt})
}

func DeleteMovie(c *gin.Context) {
	//var movieId MovieId
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "#"})
		return
	}

	if err = db.GetDb().Delete(&models.Movie{}, id).Error; err != nil {
		log.Printf("error to delete movie by id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "movie deleted"})
}

func GetMovieById(c *gin.Context) {
	var movie models.Movie

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "#"})
		return
	}

	if err = db.GetDb().Where("id = ?", id).First(&movie).Error; err != nil {
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//	// Handle record not found error...
		//}
		log.Printf("error to get movie by id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": movie})
}
