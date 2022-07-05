package serializer

type GetAllSeatsSerializer struct {
	TheaterID int `json:"theater" binding:"required"`
}
