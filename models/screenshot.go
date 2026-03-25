package models

type Screenshot struct {
	ID       string `gorm:"primaryKey;type:varchar(36)" json:"id"`
	GameID   string `gorm:"type:varchar(36);index" json:"game_id"`
	ImageSrc string `json:"image_src"`

	Game Game `gorm:"foreignKey:GameID;references:ID"`
}
