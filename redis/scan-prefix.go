package redis

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awesome-sphere/as-seating/serializer"
)

func extractSeatID(iterVal string) int {
	slice := strings.Split(iterVal, "-")
	if len(slice) < 3 {
		return -1
	}
	seatID, err := strconv.Atoi(slice[2])
	if err != nil {
		return -1
	}
	return seatID
}

func ScanPrefix(theaterID int64, timeSlotID int64, status string) ([]serializer.SeatOutputSerializer, error) {
	output := make([]serializer.SeatOutputSerializer, 0)
	noFilter := status == ""

	prefix := fmt.Sprintf("%d-%d-*", theaterID, timeSlotID)
	iter := CLIENT.Scan(0, prefix, 0).Iterator()
	for iter.Next() {
		if seatID := extractSeatID(iter.Val()); seatID > 0 {
			res, _ := CLIENT.Get(iter.Val()).Result()
			if noFilter || res == status {
				output = append(output, serializer.SeatOutputSerializer{
					SeatID:     seatID,
					TimeSlotID: timeSlotID,
					TheaterID:  theaterID,
					Status:     res,
				})
			}
		}
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}
	return output, nil
}
