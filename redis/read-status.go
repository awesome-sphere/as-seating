package redis

import (
	"fmt"
)

func ReadStatus(theaterID, timeslotID, seatID int64) (string, error) {
	return CLIENT.Get(fmt.Sprintf(
		"%d-%d-%d",
		theaterID,
		timeslotID,
		seatID,
	)).Result()
}
