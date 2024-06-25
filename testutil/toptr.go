package testutil

import (
	"time"
)

func ToStringPtr(s string) *string {
	return &s
}

func ToUint64Ptr(u uint64) *uint64 {
	return &u
}
func ToTimePtr(t time.Time) *time.Time {
	return &t
}
