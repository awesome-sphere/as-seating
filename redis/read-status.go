package redis

import (
	"fmt"

	"github.com/awesome-sphere/as-seating/serializer"
)

func ReadStatus(validatedInput serializer.CheckSeatInputSerializer) (string, error) {
	return CLIENT.Get(fmt.Sprintf(
		"%d-%d-%d",
		validatedInput.TheaterID,
		validatedInput.TimeSlotID,
		validatedInput.SeatID,
	)).Result()
}
