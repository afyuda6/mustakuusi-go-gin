package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mustakuusi-go-gin/models"
	"net/http"
)

type CharacterResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ImageSrc    string `json:"imageSrc"`
	Description string `json:"description"`
}

type CharactersController struct {
	DB *gorm.DB
}

// Index godoc
// @Tags Characters
// @Produce */*
// @Router /characters [get]
func (cc CharactersController) Index(c *gin.Context) {
	var characters []models.Character

	err := cc.DB.
		Order("position ASC").
		Find(&characters).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var result []CharacterResponse

	for _, ch := range characters {
		result = append(result, CharacterResponse{
			ID:          ch.ID,
			Name:        ch.Name,
			ImageSrc:    ch.ImageSrc,
			Description: ch.Description,
		})
	}

	c.JSON(http.StatusOK, result)
}
