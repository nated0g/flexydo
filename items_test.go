package main

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewItem(t *testing.T) {
	// Check end time is created properly
	goodStart := roundStart(time.Now(), minDur)
	goodEnd := roundEnd(time.Now(), minDur)

	want := item{
		name:  "testItem0",
		start: goodStart,
		end:   goodEnd,
	}
	got, _ := newItem("testItem1", goodStart, goodEnd)
	if !cmp.Equal(got.end, want.end) {
		t.Errorf("newItem, got %v, want %v", got.end, want.end)
	}

}
