package controllers

import (
	"latihan-gin/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SelectAllDay(c *gin.Context) {
	res := model.SelectAllDay()
	c.JSON(200, gin.H{
		"Message": "AllDay Succesfully",
		"data":    res,
	})
}

func SelectDay(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := model.SelectDay(id)
	c.JSON(200, gin.H{
		"Message": "Show Day Succesfully",
		"data":    res,
	})
}

func InsertDay(c *gin.Context) {
	var input model.Day
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.InsertDay(input.Name)
	c.JSON(200, gin.H{
		"message": "Create Succesfully",
	})
}

func UpdateDay(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input model.Day
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.UpdateDay(id, input.Name)
	c.JSON(200, gin.H{
		"message": "Update Succesfully",
	})
}

func DeleteDay(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	model.DeleteDay(id)
	c.JSON(200, gin.H{
		"message": "Delete Succesfully",
	})
}
