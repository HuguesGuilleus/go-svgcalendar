// BSD 3-Clause License
// Copyright (c) 2021, Hugues Guilleus
// All rights reserved.

package svgcalendar

import (
	"sort"
	"time"
)

type Calendar struct {
	Years    map[int]Year
	location *time.Location
}

// Year index by YearDay from 0 to 265 or 266.
type Year []Day

// One Day
type Day struct {
	Date  time.Time
	Value int
	Data  []interface{}
}

// New create a new calendar. if l is nil, l setted to time.Local
func New(l *time.Location) Calendar {
	if l == nil {
		l = time.Local
	}
	return Calendar{
		Years:    make(map[int]Year),
		location: l,
	}
}

// For the day, add value and additional data.
func (c *Calendar) Add(day time.Time, value int, data interface{}) {
	day = day.In(c.location).Truncate(time.Hour * 24)

	dy := day.Year()
	y := c.Years[dy]
	if y == nil {
		if (dy%4 == 0 && dy%100 != 0) || dy%400 == 0 {
			y = make([]Day, 366)
		} else {
			y = make([]Day, 365)
		}
		c.Years[day.Year()] = y
	}

	d := y[day.YearDay()-1]
	d.Date = day
	d.Value += value
	if data != nil {
		d.Data = append(d.Data, data)
	}
	y[day.YearDay()-1] = d
}

// Return the maximal value of this years.
func (y Year) max() (max int) {
	for _, d := range y {
		if d.Value > max {
			max = d.Value
		}
	}
	return
}

// Return sorted years.
func (c *Calendar) years() (years []int) {
	years = make([]int, 0, len(c.Years))
	for y := range c.Years {
		years = append(years, y)
	}
	sort.Slice(years, func(i int, j int) bool { return years[i] > years[j] })
	return
}
