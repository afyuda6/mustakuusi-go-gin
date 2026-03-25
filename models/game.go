package models

import (
	"strings"
	"time"
)

type Game struct {
	ID          string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Title       string    `json:"title"`
	ImageSrc    string    `json:"image_src"`
	ReleaseDate time.Time `json:"release_date"`
	Description string    `json:"description"`

	CategoriesRaw string `gorm:"column:categories" json:"-"`

	Detail            string `json:"detail"`
	PrivacyPolicyLink string `json:"privacy_policy_link"`
	DownloadLink      string `json:"download_link"`
	LongDescription   string `json:"long_description"`

	Screenshots    []Screenshot    `gorm:"foreignKey:GameID"`
	GameCharacters []GameCharacter `gorm:"foreignKey:GameID"`
}

func (g Game) Categories() []string {
	if g.CategoriesRaw == "" {
		return []string{}
	}

	trimmed := strings.Trim(g.CategoriesRaw, "{}")
	parts := strings.Split(trimmed, ",")

	var result []string
	for _, p := range parts {
		result = append(result, strings.Trim(p, `"`))
	}

	return result
}
