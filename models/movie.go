package models

import "time"

type MovieReq struct {
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
	Genre  string  `json:"genre"`
}

type Movie struct {
	ID        int64      `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name      string     `json:"name" gorm:"name"`
	Rating    float32    `json:"rating" gorm:"rating"`
	Genre     string     `json:"genre" gorm:"genre"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated-at"`
}

type MovieId struct {
	ID int64 `json:"id" form:"id"`
}

type MovieFilter struct {
	Name     string   `form:"name"`
	Rating   *float32 `form:"rating"`
	Genre    string   `form:"genre"`
	DateFrom *string  `form:"date_from"`
	DateTo   *string  `form:"date_to"`
	Page     *int     `form:"page"`
}

func (Movie) TableName() string {
	return "movie"
}
