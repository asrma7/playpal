package models

type Feed struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	DateTime string `json:"datetime"`
	Players  int64  `json:"players"`
	Location string `json:"location"`
}
