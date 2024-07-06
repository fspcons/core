package utils

import (
	"strings"
	"time"
)

// IsEmpty checks whether a string after space trimming is empty or not.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// AsRef converts T to pointer
func AsRef[T any](v T) *T {
	return &v
}

// Now return the current time
func Now() time.Time {
	return time.Now().UTC()
}
