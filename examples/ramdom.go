// BSD 3-Clause License
// Copyright (c) 2021, Hugues Guilleus
// All rights reserved.

package main

import (
	"crypto/rand"
	"github.com/HuguesGuilleus/go-svgcalendar"
	"log"
	"os"
	"time"
)

func main() {
	// Genrate random value
	buff := make([]byte, 366)
	rand.Read(buff)

	// Add ramdom value to the calendar
	c := svgcalendar.New(nil)
	for day := 0; day < 365; day++ {
		t := time.Date(2021, time.January, day+1, 0, 0, 1, 0, time.Local)
		c.Add(t, int(buff[day]), nil)
	}

	// Crete an output file
	out, err := os.Create("random.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Generate the SVG
	if err := c.SVG(out, nil); err != nil {
		log.Fatal(err)
	}
}
