package models

type Character struct {
	ID          string `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name        string `json:"name"`
	ImageSrc    string `gorm:"column:image_src" json:"image_src"`
	Description string `json:"description"`
	Position    int    `json:"position"`

	GameCharacters []GameCharacter `gorm:"foreignKey:CharacterID"`
}
