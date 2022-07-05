package serializer

type SeatOutputSerializer struct {
	ID         int64  `json:"id"`
	TimeSlotId int    `json:"time_slot_id"`
	TheaterId  int    `json:"theater_id"`
	SeatNumber int    `json:"seat_number"`
	BookedTime string `json:"booked_time"`
	BookedBy   int    `json:"booked_by"`
}
