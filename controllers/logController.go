package controllers

import (
	"net/http"
	"update-data/database"
	"update-data/helpers"
	"update-data/models"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func UpdateStatus(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	log := models.Log{}

	if contentType == appJSON {
		if err := c.ShouldBindBodyWithJSON(&log); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&log); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	if log.Water <= 5 {
		log.StatusWater = "aman"
	} else if log.Water >= 6 && log.Water < 8 {
		log.StatusWater = "siaga"
	} else if log.Water >= 8 {
		log.StatusWater = "bahaya"
	}

	if log.Wind <= 6 {
		log.StatusWind = "aman"
	} else if log.Wind >= 7 && log.Wind < 15 {
		log.StatusWind = "siaga"
	} else if log.Wind >= 15 {
		log.StatusWind = "bahaya"
	} 

	if err := db.Debug().Create(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"water":        &log.Water,
		"wind":         &log.Wind,
		"status_water": &log.StatusWater,
		"status_wind":  &log.StatusWind,
	})
}
