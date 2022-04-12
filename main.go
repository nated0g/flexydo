package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type item struct {
	id       int
	name     string
	start    time.Time
	end      time.Time
	duration time.Duration
	fixed    bool
	active   bool
	dead     bool
}

var (
	past   []item
	future []item
)

func newItem(name string, start time.Time, duration time.Duration) *item {
	i := item{name: name, start: start, duration: duration}
	i.end = start.Add(duration)
	return &i
}

func main() {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle(time.Now().String()).
		SetBackgroundColor(tcell.NewRGBColor(0, 172, 215))

	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
