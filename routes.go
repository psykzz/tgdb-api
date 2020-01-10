package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func handleRoutes(db *gorm.DB) *gin.Engine {

	db.Set("gorm:auto_preload", true)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Player based routes
	r.GET("/players", func(c *gin.Context) {
		players := []Player{}
		db.Preload("ConnectionLogs").Order("ckey asc").Find(&players)
		c.JSON(200, players)
	})
	r.GET("/players/:ckey", func(c *gin.Context) {
		ckey := c.Param("ckey")
		singlePlayer := Player{}
		db.Preload("ConnectionLogs").Where("Ckey = ?", ckey).Order("ckey asc").First(&singlePlayer)
		c.JSON(200, singlePlayer)
	})
	r.GET("/players/:ckey/connections", func(c *gin.Context) {
		ckey := c.Param("ckey")
		connectionLogs := []ConnectionLog{}
		db.Where("Ckey = ?", ckey).Order("ckey asc").Find(&connectionLogs)
		c.JSON(200, connectionLogs)
	})
	// Round based routes
	r.GET("/rounds/:id", func(c *gin.Context) {
		roundID := c.Param("id")
		roundDetail := Round{}
		db.First(&roundDetail, roundID)
		c.JSON(200, roundDetail)
	})
	r.GET("/rounds/:id/players", func(c *gin.Context) {
		roundID := c.Param("id")
		roundDetail := Round{}
		db.First(&roundDetail, roundID)
		c.JSON(200, roundDetail)
	})
	r.GET("/rounds", func(c *gin.Context) {
		rounds := []Round{}
		db.Order("id desc").Limit(10).Find(&rounds)
		c.JSON(200, rounds)
	})
	r.Run()

	return r
}
