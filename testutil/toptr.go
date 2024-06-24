package testutil

import "time"

func ToStringPtr(s string) *string {
	return &s
}

func ToTimePtr(t time.Time) *time.Time {
	return &t
}
