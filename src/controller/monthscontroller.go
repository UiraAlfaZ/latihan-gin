package controllers

import (
	"latihan-gin/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SelectAllMonth(c *gin.Context) {
	res := model.SelectAllMonth()
	c.JSON(200, gin.H{
		"message": "AllMonth Successful",
		"data":    res,
	})
}

func SelectMonth(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := model.SelectMonth(id)
	c.JSON(200, gin.H{
		"message": "Show Month Successful",
		"data":    res,
	})
}

func InsertMonth(c *gin.Context) {
	var input model.Month
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.InsertMonth(input.Name, input.Day)
	c.JSON(200, gin.H{
		"message": "Created Successfully",
	})
}

func UpdateMonth(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input model.Month
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.UpdateMonth(id, input.Name, input.Day)
	c.JSON(200, gin.H{
		"message": "Updated Successfully",
	})
}

func DeleteMonth(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	model.DeleteMonth(id)
	c.JSON(200, gin.H{
		"message": "Deleted Successfully",
	})
}
