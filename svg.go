// BSD 3-Clause License
// Copyright (c) 2021, Hugues Guilleus
// All rights reserved.

package svgcalendar

import (
	_ "embed"
	"encoding/xml"
	"io"
	"strconv"
	"time"
)

// Theme is used to print SVG.
type Theme struct {
	// Title text style
	Title ThemeText

	// The number of lines for hover information. If zero, no display.
	HoverLines int
	// The text style for hover information
	Hover ThemeText

	// The color pallete,
	Colors []string
	// The width and height of one day.
	Length int
	// The round of day
	Round int
	// The space between two day
	Space int
}

// All style theme for text.
type ThemeText struct {
	// The CSS font propertys
	Style string
	// The height
	Height int
	// The margin bottom
	Bottom int
}

var Default = &Theme{
	Title: ThemeText{
		Style:  `font-family:sans;font-size:30px;`,
		Height: 50,
		Bottom: 20,
	},
	HoverLines: 10,
	Hover: ThemeText{
		Style:  `font-family:sans;font-size:20px;`,
		Height: 20,
		Bottom: 10,
	},
	Colors: []string{
		"hsl(0, 0%, 90%)",
		"hsl(210, 100%, 90%)",
		"hsl(210, 100%, 80%)",
		"hsl(210, 100%, 60%)",
		"hsl(210, 100%, 20%)",
	},
	Length: 20,
	Round:  5,
	Space:  3,
}

//go:embed hover.js
var hover string

type text struct {
	Y     int    `xml:"y,attr"`
	Style string `xml:"style,attr"`
	Text  int    `xml:",chardata"`
}

type rect struct {
	X      int         `xml:"x,attr"`
	Y      int         `xml:"y,attr"`
	Width  int         `xml:"width,attr"`
	Height int         `xml:"height,attr"`
	Rx     int         `xml:"rx,attr"`
	Ry     int         `xml:"ry,attr"`
	Fill   string      `xml:"fill,attr"`
	Data   interface{} `xml:"data,omitempty"`
	Date   time.Time   `xml:"date,attr"`
	Value  int         `xml:"value,attr"`
}

type script struct {
	Content string `xml:",cdata"`
}

// Print SVG graphics. If t in nil, SVG use Default theme.
func (c *Calendar) SVG(w io.Writer, t *Theme) error {
	if t == nil {
		t = Default
	}
	size := t.Length + t.Space // the size of one cell

	// Print headers
	w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?><svg xmlns="http://www.w3.org/2000/svg" version="1.1" viewBox="0 0 ` +
		strconv.Itoa(size*53-t.Space) +
		` ` +
		strconv.Itoa(
			len(c.Years)*(t.Title.Height+t.Title.Bottom+size*7)+
				t.HoverLines*t.Hover.Height) +
		`" >`))
	defer w.Write([]byte(`</svg>`))

	enc := xml.NewEncoder(w)
	height := 0
	for _, y := range c.years() {
		// Print the title
		yc := c.Years[y]
		height += t.Title.Height
		err := enc.Encode(&text{
			Y:     height,
			Style: t.Title.Style,
			Text:  y,
		})
		if err != nil {
			return err
		}
		height += t.Title.Bottom

		// Print cells
		max := yc.max()
		offset := int(time.Date(y, time.January, 1, 0, 0, 0, 0, c.location).Weekday())
		offset = (offset+6)%7 - 1
		for i, d := range yc {
			if d.Date.IsZero() {
				d.Date = time.Date(y, time.January, i+1, 0, 0, 0, 0, c.location)
			}
			fill := t.Colors[0]
			if d.Value != 0 {
				fill = t.Colors[d.Value*(len(t.Colors)-2)/max+1]
			}
			err = enc.Encode(&rect{
				X:      (d.Date.YearDay() + offset) / 7 * size,
				Y:      height + int(d.Date.Weekday()+6)%7*size,
				Width:  t.Length,
				Height: t.Length,
				Rx:     t.Round,
				Ry:     t.Round,
				Fill:   fill,
				Data:   d.Data,
				Date:   d.Date,
				Value:  d.Value,
			})
			if err != nil {
				return err
			}
		}
		height += size * 7
	}

	if t.HoverLines != 0 {
		height += t.Hover.Height
		enc.Encode(&text{
			Y:     height,
			Style: t.Hover.Style,
		})
		enc.Encode(&script{Content: hover})
	}

	return nil
}
