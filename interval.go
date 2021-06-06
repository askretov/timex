package timex

import (
	"fmt"
	"time"
)

// Interval describes a time interval between start and end time values
type Interval struct {
	start time.Time
	end   time.Time
}

// NewInterval returns a new instance of Interval between start and end
// Input is order-independent, the smaller value will be used as the start when bigger as the end
func NewInterval(t1, t2 time.Time) Interval {
	i := Interval{start: t1, end: t2}
	if t2.Before(t1) {
		i = Interval{start: t2, end: t1}
	}
	return i
}

// Start returns i's start value
func (i Interval) Start() time.Time {
	return i.start
}

// End returns i's end value
func (i Interval) End() time.Time {
	return i.end
}

// Days returns a duration of interval in days
func (i Interval) Days() float64 {
	return float64(i.Nanoseconds()) / float64(Day)
}

// Hours returns a duration of interval in hours
func (i Interval) Hours() float64 {
	return float64(i.Nanoseconds()) / float64(Hour)
}

// Minutes returns a duration of interval in minutes
func (i Interval) Minutes() float64 {
	return float64(i.Nanoseconds()) / float64(Minute)
}

// Seconds returns a duration of interval in seconds
func (i Interval) Seconds() float64 {
	return float64(i.Nanoseconds()) / float64(Second)
}

// Milliseconds returns a duration of interval in milliseconds
func (i Interval) Milliseconds() float64 {
	return float64(i.Nanoseconds()) / float64(Millisecond)
}

// Microseconds returns a duration of interval in microseconds
func (i Interval) Microseconds() float64 {
	return float64(i.Nanoseconds()) / float64(Microsecond)
}

// Nanoseconds returns a duration of interval in hours
func (i Interval) Nanoseconds() int64 {
	return i.end.Sub(i.start).Nanoseconds()
}

// IsZero reports whether i start & end values are zero time values
func (i Interval) IsZero() bool {
	return i.start.IsZero() && i.end.IsZero()
}

// Contains reports whether t is within of i (closed interval strategy, start <= t <= end)
func (i Interval) Contains(t time.Time) bool {
	return t.UnixNano() >= i.start.UnixNano() && t.UnixNano() <= i.end.UnixNano()
}

// Duration returns i's duration as time.Duration value
func (i Interval) Duration() time.Duration {
	return time.Duration(i.Nanoseconds())
}

// String returns string representation of i
func (i Interval) String() string {
	return fmt.Sprintf("%s - %s", i.start.Format(time.RFC3339), i.end.Format(time.RFC3339))
}

// StringDates returns string representation of i as date part only
func (i Interval) StringDates() string {
	return fmt.Sprintf("%s - %s", i.start.Format("2006-01-02"), i.end.Format("2006-01-02"))
}

// HalfOpenEnd returns i as a half-open interval where the open side is on the end [start, end)
// Actually, it returns copy of i where end = end - 1 nanosecond
func (i Interval) HalfOpenEnd() Interval {
	return NewInterval(i.Start(), i.End().Add(-1))
}
