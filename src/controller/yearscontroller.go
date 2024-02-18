package controllers

import (
	"latihan-gin/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SelectAllYear(c *gin.Context) {
	res := model.SelectAllYear()
	c.JSON(200, gin.H{
		"Message": "AllYear Succesfully",
		"data":    res,
	})
}

func SelectYear(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := model.SelectYear(id)
	c.JSON(200, gin.H{
		"Message": "Show Year Succesfully",
		"data":    res,
	})
}

func InsertYear(c *gin.Context) {
	var input model.Year
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.InsertYear(input.Name)
	c.JSON(200, gin.H{
		"message": "Create Succesfully",
	})
}

func UpdateYear(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input model.Year
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.UpdateYear(id, input.Name)
	c.JSON(200, gin.H{
		"message": "Update Succesfully",
	})
}

func DeleteYear(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	model.DeleteYear(id)
	c.JSON(200, gin.H{
		"message": "Delete Succesfully",
	})
}
