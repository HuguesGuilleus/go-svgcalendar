# go-svgcalendar
[![Go Reference](https://pkg.go.dev/badge/github.com/HuguesGuilleus/go-svgcalendar.svg)](https://pkg.go.dev/github.com/HuguesGuilleus/go-svgcalendar)

Create a SVG calendar with day value.


## Import
```sh
go get github.com/HuguesGuilleus/go-svgcalendar
```

## Usage

```go

package main

import (
	"github.com/HuguesGuilleus/go-svgcalendar"
	"time"
	"io"
)

func generation() {
	// Create teh calendar
	c := svgcalendar.New(nil)

	// Add value
	t := time.Now()
	v := 42
	additionalInformations := nil
	c.Add(t, v, additionalInformations)

	// Generate and save the SVG
	var out io.Writer
	c.SVG(out, nil)
}
```

See the file `examples/random.go` to a complete example.
