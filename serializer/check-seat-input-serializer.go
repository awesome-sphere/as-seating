package serializer

type CheckSeatInputSerializer struct {
	TheaterID  int64 `json:"theater_id" binding:"required"`
	SeatID     int64 `json:"seat_id" binding:"required"`
	TimeSlotID int64 `json:"time_slot_id" binding:"required"`
}
