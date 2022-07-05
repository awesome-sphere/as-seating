package serializer

type SeatsInputSerializer struct {
	TheaterID int `json:"theater" binding:"required"`
}
