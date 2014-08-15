package lib

import (
	"fmt"
	"time"
)

func Generate() {

}

func CreateVersion() string {
	now := time.Now()
	parts := []int{
		now.Year(),
		int(now.Month()),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	}
	stamp := ""
	for _, part := range parts {
		stamp += fmt.Sprintf("%02d", part)
	}

	return stamp
}
