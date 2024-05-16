package logger

import (
	"math"
	"time"
	"fmt"
)

func (c *Club)calculateRevenue(event Event, tableInd int) {
	startTime := c.Tables[tableInd].Start
	endTime := event.Time

	// Calculate the duration in minutes
	duration := endTime.Sub(startTime)
	c.Tables[tableInd].Duration += duration
	minutesWorked := duration.Minutes()

	// Calculate the number of hours, rounding up if there are remaining minutes
	hoursWorked := int(minutesWorked / 60)
	if math.Mod(minutesWorked, 60) > 0 {
		hoursWorked++
	}

	c.Tables[tableInd].Income += hoursWorked * c.HourlyRate
}

func fmtDuration(d time.Duration) string {
    d = d.Round(time.Minute)
    h := d / time.Hour
    d -= h * time.Hour
    m := d / time.Minute
    return fmt.Sprintf("%02d:%02d", h, m)
}