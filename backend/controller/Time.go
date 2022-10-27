package controller

import (
	"github.com/phanukorn2644/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Time
func CreateTime(c *gin.Context) {
	var time entity.Time
	if err := c.ShouldBindJSON(&time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&time).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": time})
}

// GET /Time/:id
func GetTime(c *gin.Context) {
	var time entity.Time
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time})
}

// GET /Time
func ListTimes(c *gin.Context) {
	var times []entity.Time
	if err := entity.DB().Raw("SELECT * FROM times").Scan(&times).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": times})
}

// DELETE /Time/:id
func DeleteTime(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM times WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Time
func UpdateTime(c *gin.Context) {
	var time entity.Time
	if err := c.ShouldBindJSON(&time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", time.ID).First(&time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	if err := entity.DB().Save(&time).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time})
}
