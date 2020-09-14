package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Truck struct {
	Gear  uint   `json:"gear"`
	Id    uint   `json:"id"`
	Model string `json:"model"`
}

func getAllTrucks(c *gin.Context) {
	var truck_from_db []Truck
	err := db.Find(&truck_from_db).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, truck_from_db)
	}
}

func getTruck(c *gin.Context) {
	var truck Truck
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&truck).Error
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)

	} else {
		c.JSON(200, truck)
	}
}
func addTruck(c *gin.Context) {
	var truck Truck
	c.BindJSON(&truck)
	db.Create(&truck)
	c.JSON(200, truck)
	db.Save(&truck)

}
func updateTruck(c *gin.Context) {
	var truck Truck
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&truck).Error
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)

	} else {
		c.BindJSON(&truck)
		db.Save(&truck)
		c.JSON(200, truck)
	}
}
func deleteTruck(c *gin.Context) {
	id := c.Params.ByName("id")
	var truck Truck
	err := db.Where("id  = ? ", id).Delete(&truck).Error
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"Deleted": "Deleted with id " + string(id)})
	}
}
func main() {
	//
	//
	//
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatal("error")
	}
	defer db.Close()
	db.AutoMigrate(&Truck{})
	r := gin.Default()
	r.GET("/", getAllTrucks)
	r.GET("/trucks/:id", getTruck)
	r.POST("/addingtruck", addTruck)
	r.PUT("/updatetruck/:id", updateTruck)
	r.DELETE("/delete/:id", deleteTruck)
	r.Run(":8000")
}
