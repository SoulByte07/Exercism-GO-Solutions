package clock

import "fmt"

type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	return fromTotalMinutes(h*60 + m)
}

func (c Clock) Add(m int) Clock {
	return fromTotalMinutes(c.h*60 + c.m + m)
}

func (c Clock) Subtract(m int) Clock {
	return fromTotalMinutes(c.h*60 + c.m - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func fromTotalMinutes(total int) Clock {
	const minutesPerDay = 24 * 60

	total %= minutesPerDay
	if total < 0 {
		total += minutesPerDay
	}

	return Clock{
		h: total / 60,
		m: total % 60,
	}
}
