package redis

import (
	"log"

	"github.com/awesome-sphere/as-seating/utils"
)

func initSeats() {

	for theater := 1; theater <= 5; theater++ {
		theater := int64(theater)
		for timeSlot := 1; timeSlot <= 5*5*5; timeSlot++ {
			timeSlot := (theater-1)*125 + int64(timeSlot)
			for seat := 1; seat <= utils.TertiaryOperator(timeSlot == 263, 20, 40).(int); seat++ {
				seatStatus, err := ReadStatus(theater, timeSlot, int64(seat))
				if err == nil && seatStatus != AVAILABLE {
					UpdateStatus(theater, timeSlot, seat, AVAILABLE, seat == 1 && timeSlot%50 == 1)
				}
			}
		}
	}

	log.Printf("Initialized seats.")
}
