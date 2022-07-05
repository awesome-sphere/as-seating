package models

import (
	"time"
)

type TimeSlot struct {
	// gorm.Model
	Id        int64     `json:"id" gorm:"primaryKey;autoincrement;not null"`
	MovieId   int       `json:"movie_id" gnorm:"not null"`
	Time      time.Time `json:"time" gnorm:"not null"`
	TheaterId int       `json:"theater_id" gnorm:"not null"`
	Theater   Theater   `gorm:"foreignKey:theater_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TimeSlots_1 struct {
	TimeSlot
}

type TimeSlots_2 struct {
	TimeSlot
}

type TimeSlots_3 struct {
	TimeSlot
}
type TimeSlots_4 struct {
	TimeSlot
}
type TimeSlots_5 struct {
	TimeSlot
}
