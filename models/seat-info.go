package models

import "time"

type SeatStatus string

const (
	Available SeatStatus = "AVAILABLE"
	Booked    SeatStatus = "BOOKED"
	Bought    SeatStatus = "BOUGHT"
)

type SeatInfo struct {
	Id         int64      `json:"id" gorm:"primaryKey;autoincrement;"`
	TimeSlotId int        `json:"time_slot_id"`
	TimeSlot   TimeSlot   `gorm:"foreignKey:time_slot_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TheaterId  int        `json:"theater_id" gnorm:"not null"`
	Theater    Theater    `gorm:"foreignKey:theater_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SeatNumber int        `json:"seat_number" gorm:"not null"`
	Status     SeatStatus `json:"status" gorm:"default:'AVAILABLE'" sql:"type:seat_status"`
	BookedTime time.Time  `json:"booked_time" gnorm:"default:time.now();not null"`
	BookedBy   int        `json:"booked_by" gorm:"not null"`
}

type SeatInfos_1 struct {
	SeatInfo
	TimeSlot TimeSlots_1 `gorm:"foreignKey:time_slot_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SeatInfos_2 struct {
	SeatInfo
	TimeSlot TimeSlots_2 `gorm:"foreignKey:time_slot_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SeatInfos_3 struct {
	SeatInfo
	TimeSlot TimeSlots_3 `gorm:"foreignKey:time_slot_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SeatInfos_4 struct {
	SeatInfo
	TimeSlot TimeSlots_3 `gorm:"foreignKey:time_slot_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SeatInfos_5 struct {
	SeatInfo
	TimeSlot TimeSlots_4 `gorm:"foreignKey:time_slot_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
