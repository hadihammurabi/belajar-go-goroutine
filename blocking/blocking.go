package blocking

import (
	"time"
)

func Forever() {
	for {
		<-time.After(5 * time.Minute)
	}
}
