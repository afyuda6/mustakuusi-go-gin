package models

type GameCharacter struct {
	GameID      string `gorm:"type:varchar(36);index" json:"game_id"`
	CharacterID string `gorm:"type:varchar(36);index" json:"character_id"`

	Game      Game      `gorm:"foreignKey:GameID;references:ID"`
	Character Character `gorm:"foreignKey:CharacterID;references:ID"`
}
