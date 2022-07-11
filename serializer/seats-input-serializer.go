package serializer

type SeatsInputSerializer struct {
	TheaterID  int64 `json:"theater_id" binding:"required"`
	TimeSlotID int64 `json:"time_slot_id" binding:"required"`
}
