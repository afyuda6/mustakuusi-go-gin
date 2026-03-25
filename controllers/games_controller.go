package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mustakuusi-go-gin/models"
	"net/http"
	"sort"
	"time"
)

type GameResponse struct {
	ID                string   `json:"id"`
	Title             string   `json:"title"`
	ImageSrc          string   `json:"imageSrc"`
	Date              string   `json:"date"`
	Description       string   `json:"description"`
	Categories        []string `json:"categories"`
	Detail            string   `json:"detail"`
	PrivacyPolicyLink string   `json:"privacyPolicyLink"`
	DownloadLink      string   `json:"downloadLink"`
	LongDescription   string   `json:"longDescription"`
	Screenshots       []string `json:"screenshots"`
	Characters        []string `json:"characters"`
}

type GamesController struct {
	DB *gorm.DB
}

// Index godoc
// @Tags Games
// @Produce */*
// @Router /games [get]
func (gc GamesController) Index(c *gin.Context) {
	var games []models.Game

	err := gc.DB.
		Preload("Screenshots").
		Preload("GameCharacters.Character").
		Order("release_date DESC").
		Order("id DESC").
		Find(&games).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")

	var result []GameResponse

	for _, game := range games {

		sort.Slice(game.Screenshots, func(i, j int) bool {
			return game.Screenshots[i].ID < game.Screenshots[j].ID
		})

		screenshots := []string{}
		for _, s := range game.Screenshots {
			screenshots = append(screenshots, s.ImageSrc)
		}

		sort.Slice(game.GameCharacters, func(i, j int) bool {
			return game.GameCharacters[i].Character.Position <
				game.GameCharacters[j].Character.Position
		})

		characters := []string{}
		for _, gc := range game.GameCharacters {
			characters = append(characters, gc.Character.ID)
		}

		result = append(result, GameResponse{
			ID:                game.ID,
			Title:             game.Title,
			ImageSrc:          game.ImageSrc,
			Date:              game.ReleaseDate.In(loc).Format(time.RFC3339),
			Description:       game.Description,
			Categories:        game.Categories(),
			Detail:            game.Detail,
			PrivacyPolicyLink: game.PrivacyPolicyLink,
			DownloadLink:      game.DownloadLink,
			LongDescription:   game.LongDescription,
			Screenshots:       screenshots,
			Characters:        characters,
		})
	}

	c.JSON(http.StatusOK, result)
}
