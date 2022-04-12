package main

import (
	"fmt"
	"time"
)

type item struct {
	name   string
	start  time.Time
	end    time.Time
	fixed  bool
	active bool
	dead   bool
	prev   *item
	next   *item
}

type list struct {
	len  int
	head *item
	tail *item
}

var (
	past   list
	future list
)

const (
	Day time.Duration = 24 * time.Hour
)

const minDur = "30m"

func roundStart(t time.Time, durStr string) time.Time {
	d, _ := time.ParseDuration(durStr)
	return t.Truncate(d)
}

func midnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func endOfDay(t time.Time) time.Time {
	return midnight(t).Add(Day)
}

func roundEnd(t time.Time, durStr string) time.Time {
	d, _ := time.ParseDuration(durStr)
	newT := t.Truncate(d)
	if newT != t {
		newT.Add(d)
	}
	return newT
}

func newItem(name string, start time.Time, end time.Time) (*item, error) {

	start = roundStart(start, minDur)
	end = roundEnd(end, minDur)

	i := item{name: name, start: start, end: end}
	return &i, nil
}

// Create a new list with one item, which spans from now to the end of the day
func initFuture() list {
	l := list{}
	i, _ := newItem("", time.Now(), endOfDay(time.Now()))
	l.head = i
	l.tail = i
	l.len = 1
	return l

}

func (l *list) prepend(i *item) {
	l.head.start = i.end
	i.next = l.head
	l.head.prev = i
	l.head = i
	l.len++
}

func initPast() list {
	return list{}
}

func main() {
	f := initFuture()

	i, _ := newItem("test", time.Now(), time.Now().Add(time.Hour*2))
	f.prepend(i)

	i, _ = newItem("test", time.Now(), time.Now().Add(time.Hour*1))
	f.prepend(i)

	fmt.Printf("%+v\n", f)
	fmt.Printf("%+v\n", f.head.start)
	fmt.Printf("%+v\n", f.head.end)

	fmt.Printf("%+v\n", f.head.next.start)
	fmt.Printf("%+v\n", f.head.next.end)

	fmt.Printf("%+v\n", f.tail.start)
	fmt.Printf("%+v\n", f.tail.end)

}
