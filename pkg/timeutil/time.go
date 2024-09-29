package timeutil

import (
	"log"
	"time"
	_ "time/tzdata"
)

var VNTLocation *time.Location

func init() {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Fatalf("Load location failed %v", err)
	}

	VNTLocation = location
}

func NowInVNT() time.Time {
	return time.Now().In(VNTLocation)
}

func FindNextDays(current time.Time, days int) time.Time {
	return current.AddDate(0, 0, days)
}

func ShiftDays(current time.Time, days int) time.Time {
	return current.AddDate(0, 0, days)
}
