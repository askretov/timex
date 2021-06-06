package timex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTmp(tt *testing.T) {
	// Create an interval
	start := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	i := NewInterval(start, start.AddDate(0, 0, 9))
	fmt.Println(i)
	// > 2021-01-01T00:00:00Z - 2021-01-10T00:00:00Z

	// Print the interval duration in days
	fmt.Println(i.Days())
	// > 9

	// Check if t inside of i
	t := start.AddDate(0, 0, 5)
	fmt.Println(i.Contains(t))
	// > true

	// Get a half-open interval for i
	ho := i.HalfOpenEnd()
	fmt.Println(ho)
	// > 2021-01-01T00:00:00Z - 2021-01-09T23:59:59Z
}

func TestNewInterval(t *testing.T) {
	time.Now().AddDate(0, 0, 10)
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want Interval
	}{
		{
			name: "t1 < t2",
			args: args{
				t1: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			},
			want: Interval{
				start: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "t2 < t1",
			args: args{
				t1: time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			want: Interval{
				start: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "t1 = t2",
			args: args{
				t1: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			want: Interval{
				start: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInterval(tt.args.t1, tt.args.t2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func testIntervalPeriod(t *testing.T, u units) {
	tests := []struct {
		name     string
		interval Interval
		want     float64
	}{
		{
			name: "4 days",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			),
			want: 4 * float64(Day),
		},
		{
			name: "4 days",
			interval: NewInterval(
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
			),
			want: 4 * float64(Day),
		},
		{
			name: "4.5 days",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 12, 0, 0, 0, time.UTC),
			),
			want: 4.5 * float64(Day),
		},
		{
			name: "0",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
			),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got float64
			switch u {
			case Nanosecond:
				got = float64(tt.interval.Nanoseconds())
				//want = tt.want
			case Microsecond:
				got = tt.interval.Microseconds() * float64(Microsecond)
			case Millisecond:
				got = tt.interval.Milliseconds() * float64(Millisecond)
			case Second:
				got = tt.interval.Seconds() * float64(Second)
			case Minute:
				got = tt.interval.Minutes() * float64(Minute)
			case Hour:
				got = tt.interval.Hours() * float64(Hour)
			case Day:
				got = tt.interval.Days() * float64(Day)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterval_Days(t *testing.T) {
	testIntervalPeriod(t, Day)
}

func TestInterval_Hours(t *testing.T) {
	testIntervalPeriod(t, Hour)
}

func TestInterval_Minutes(t *testing.T) {
	testIntervalPeriod(t, Minute)
}

func TestInterval_Seconds(t *testing.T) {
	testIntervalPeriod(t, Second)
}

func TestInterval_Milliseconds(t *testing.T) {
	testIntervalPeriod(t, Millisecond)
}

func TestInterval_Microseconds(t *testing.T) {
	testIntervalPeriod(t, Microsecond)
}

func TestInterval_Nanoseconds(t *testing.T) {
	testIntervalPeriod(t, Nanosecond)
}

func TestInterval_IsZero(t *testing.T) {
	tests := []struct {
		name     string
		interval Interval
		want     bool
	}{
		{
			name: "both are not zero",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			),
			want: false,
		},
		{
			name:     "start is zero",
			interval: NewInterval(time.Time{}, time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC)),
			want:     false,
		},
		{
			name:     "end is zero",
			interval: NewInterval(time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC), time.Time{}),
			want:     false,
		},
		{
			name:     "both are zero",
			interval: Interval{},
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.interval.IsZero()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterval_Start(t *testing.T) {
	type fields struct {
		start time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "ok",
			fields: fields{
				start: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interval{start: tt.fields.start}
			got := i.Start()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterval_End(t *testing.T) {
	type fields struct {
		end time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "ok",
			fields: fields{
				end: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interval{end: tt.fields.end}
			got := i.End()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterval_Contains(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name     string
		interval Interval
		args     args
		want     bool
	}{
		{
			name: "t = start",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			),
			args: args{t: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "t < start, t's tz +0100",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			),
			args: args{t: time.Date(2021, 6, 1, 0, 0, 0, 0, time.FixedZone("+0100", int(time.Hour/time.Second)))},
			want: false,
		},
		{
			name: "t > start, t's tz -0100",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			),
			args: args{t: time.Date(2021, 6, 1, 0, 0, 0, 0, time.FixedZone("-0100", -int(time.Hour/time.Second)))},
			want: true,
		},
		{
			name: "t = end",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			),
			args: args{t: time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "start < t < end",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			),
			args: args{t: time.Date(2021, 6, 3, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.interval.Contains(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterval_Duration(t *testing.T) {
	tests := []struct {
		name     string
		interval Interval
		want     time.Duration
	}{
		{
			name: "4d5h15s",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 5, 0, 15, 0, time.UTC),
			),
			want: time.Duration(Day*4 + Hour*5 + Second*15),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.interval.Duration()
			fmt.Println(got)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterval_String(t *testing.T) {
	tests := []struct {
		name     string
		interval Interval
		want     string
	}{
		{
			name: "2021-06-01T00:00:00Z - 2021-06-05T05:00:15+08:00",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 5, 0, 15, 0, time.FixedZone("Asia/Singapore", int((time.Hour*8)/time.Second))),
			),
			want: "2021-06-01T00:00:00Z - 2021-06-05T05:00:15+08:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.interval.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterval_StringDates(t *testing.T) {
	tests := []struct {
		name     string
		interval Interval
		want     string
	}{
		{
			name: "2021-06-01 - 2021-06-05",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 5, 0, 15, 0, time.UTC),
			),
			want: "2021-06-01 - 2021-06-05",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.interval.StringDates()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterval_HalfOpenEnd(t *testing.T) {
	tests := []struct {
		name     string
		interval Interval
		want     Interval
	}{
		{
			name: "2021-06-01 - 2021-06-04",
			interval: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 5, 0, 0, 0, 0, time.UTC),
			),
			want: NewInterval(
				time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2021, 6, 4, 23, 59, 59, int(time.Second)-1, time.UTC),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.interval.HalfOpenEnd()
			assert.Equal(t, tt.want, got)
		})
	}
}
