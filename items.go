package main

import (
	"fmt"
	"time"

	"github.com/rwxrob/structs/qstack"
)

type item struct {
	name   string
	start  time.Time
	end    time.Time
	fixed  bool
	active bool
	dead   bool
}

type list struct {
	len  int
	head *item
	tail *item
}

type TimeList struct {
	*qstack.QS[item]
}

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
func initFuture() *TimeList {
	t := TimeList{}
	t.QS = qstack.New[item]()
	i, _ := newItem("", time.Now(), endOfDay(time.Now()))
	t.QS.Push(*i)
	return &t
}

func initPast() list {
	return list{}
}

func (t TimeList) append(i *item) {
	t.QS.Push(*i)
	top := t.QS.Peek()
	top.start = i.end
}

func main() {

	f := initFuture()
	i, _ := newItem("test", time.Now(), time.Now().Add(time.Hour*2))
	fmt.Printf("%+v\n", f.Items())
	f.append(i)
	fmt.Printf("%+v\n", f.Items())
	fmt.Printf("%+v\n", f.Current())
	f.Scan()
	fmt.Printf("%+v\n", f.Current())
	f.Back()
	fmt.Printf("%+v\n", f.Current())
	f.Back()
	fmt.Printf("%+v\n", f.Current())
	// i, _ = newItem("test2", time.Now(), time.Now().Add(time.Hour*1))
	// future.Push(*i)

	// fmt.Printf("%+v\n", future.Pop().start)
	// fmt.Printf("%+v\n", future.Pop().start)
	// fmt.Printf("%+v\n", future.Pop().start)

}
