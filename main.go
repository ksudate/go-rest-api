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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/splints", func(c *gin.Context) {
		splints := []Splint{}
		db.Find(&splints)
		c.JSON(http.StatusOK, splints)
	})
	r.Run()
}
