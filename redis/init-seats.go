package redis

import (
	"log"
)

func initSeats() {
	for theater := 1; theater <= 5; theater++ {
		theater := int64(theater)
		for timeSlot := 1; timeSlot <= 5*5*5; timeSlot++ {
			for seat := 1; seat <= 55; seat++ {
				seat = int(theater-1)*625 + int(seat)
				timeSlot := (theater-1)*125 + int64(timeSlot)
				seatStatus, err := ReadStatus(theater, timeSlot, int64(seat))
				if err == nil && seatStatus != AVAILABLE {
					UpdateStatus(theater, timeSlot, int(seat), AVAILABLE, seat == 1 && timeSlot%50 == 1)
				}
			}
		}
	}

	log.Printf("Initialized seats.")
}
