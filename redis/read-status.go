package redis

import (
	"fmt"

	"github.com/awesome-sphere/as-seating/serializer"
)

func ReadStatus(validated_input serializer.CheckSeatInputSerializer) (string, error) {
	return CLIENT.Get(fmt.Sprintf(
		"%d-%d-%d",
		validated_input.TheaterID,
		validated_input.TimeSlotID,
		validated_input.SeatID,
	)).Result()
}
