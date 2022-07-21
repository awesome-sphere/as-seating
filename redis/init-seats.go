package redis

import (
	"log"
)

func initSeats() {

	for theater := 1; theater <= 5; theater++ {
		theater := int64(theater)
		for timeSlot := (125 * (theater - 1)) + 1; timeSlot <= (125*theater-1)+(5*5*5); timeSlot++ {
			timeSlot := int64(timeSlot)
			// for seat := 1; seat <= utils.TertiaryOperator(timeSlot == 263, 20, 40).(int); seat++ {
			for seat := 1; seat <= 55; seat++ {
				seatStatus, err := ReadStatus(theater, timeSlot, int64(seat))
				if err != nil || seatStatus != AVAILABLE {
					UpdateStatus(theater, timeSlot, seat, AVAILABLE, seat == 1 && timeSlot%50 == 1)
				}
			}
		}
	}

	log.Printf("Initialized seats.")
}
