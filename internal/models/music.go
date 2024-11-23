package models

import "time"

type Music struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
    Group       string    `json:"group" gorm:"column:group_name;not null"`
	Title       string    `json:"title" gorm:"not null"`
	ReleaseDate time.Time `json:"releaseDate" gorm:"type:date"`
	Text        string    `json:"text" gorm:"type:text"`
	Link        string    `json:"link"`
	
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}