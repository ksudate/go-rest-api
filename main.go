package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Splint expect
type Splint struct {
	ID        int       `gorm:"primary_key"`
	SpNumber  int       `gorm:"type:int(11);"`
	Content   string    `gorm:"type:varchar(255);"`
	Kpt       string    `gorm:"type:varchar(255);"`
	LineID    string    `gorm:"type:varchar(255);"`
	CreatedAt time.Time `gorm:"type:datetime(6);"`
	UpdatedAt time.Time `gorm:"type:datetime(6);"`
}

func main() {
	db := ConnectDB()
	defer db.Close()
	router := gin.Default()
	// Create
	router.POST("/splint", func(c *gin.Context) {
		splint := Splint{}
		err := c.BindJSON(&splint)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		db.NewRecord(splint)
		db.Create(&splint)
		if db.NewRecord(splint) == false {
			c.JSON(http.StatusOK, splint)
		}
	})
	// Read
	router.GET("/splints", func(c *gin.Context) {
		splints := []Splint{}
		db.Find(&splints)
		c.JSON(http.StatusOK, splints)
	})
	router.GET("/splint/:id", func(c *gin.Context) {
		splint := Splint{}
		id := c.Param("id")
		db.Where("ID = ?", id).Find(&splint)
		c.JSON(http.StatusOK, splint)
	})
	// Update
	router.POST("/splint/:id", func(c *gin.Context) {
		splint := Splint{}
		id := c.Param("id")
		data := Splint{}
		err := c.BindJSON(&data)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		db.Where("ID = ?", id).First(&splint).Update(&data)
	})
	// Delete
	router.DELETE("/splint/:id", func(c *gin.Context) {
		splint := Splint{}
		id := c.Param("id")
		db.Where("ID = ?", id).Delete(&splint)
	})
	router.Run()
}
