package main

import (
	"snippetbox-go.casteleira.xyz/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2025, time.January, 1, 12, 15, 0, 0, time.UTC),
			want: "01 Jan 2025 at 12:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2025, time.January, 1, 12, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "01 Jan 2025 at 11:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.want)
		})
	}
}
