package util

import (
	"strings"
	"time"
)

// Tells whether the number has the nth bit set.
//
// bitPosition must be in the range 0-7.
func BitIsSet(number, bitPosition uint8) bool {
	return (number & (1 << bitPosition)) > 0
}

// Returns a new number with the nth bit set or clear.
//
// bitPosition must be in the range 0-7.
func BitSet(number, bitPosition uint8, doSet bool) uint8 {
	if doSet {
		return number | (1 << bitPosition)
	} else {
		return number &^ (1 << bitPosition)
	}
}

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// Returns first value if condition is true, otherwise returns second.
//
// Return type must be cast accordingly.
func Iif(cond bool, ifTrue, ifFalse interface{}) interface{} {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

// "&He && she" becomes "He & she".
func RemoveAccelAmpersands(text string) string {
	runes := []rune(text)
	buf := strings.Builder{}
	buf.Grow(len(runes)) // prealloc for performance

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == '&' && runes[i+1] != '&' {
			continue
		}
		buf.WriteRune(runes[i])
	}
	if runes[len(runes)-1] != '&' {
		buf.WriteRune(runes[len(runes)-1])
	}
	return buf.String()
}

//------------------------------------------------------------------------------

// Assembles an uint64 from two uint32.
func Make64(lo, hi uint32) uint64 {
	return (uint64(lo) & 0xffff_ffff) | ((uint64(hi) & 0xffff_ffff) << 32)
}

// Breaks an uint64 into low and high uint32.
func Break64(val uint64) (lo, hi uint32) {
	return uint32(val & 0xffff_ffff), uint32(val >> 32 & 0xffff_ffff)
}

//------------------------------------------------------------------------------

// Converts time.Duration to 100 nanoseconds.
func DurationToNano100(duration time.Duration) int64 {
	return int64(duration) * 10_000 / int64(time.Millisecond)
}

// Converts 100 nanoseconds to time.Duration.
func Nano100ToDuration(nanosec100 int64) time.Duration {
	return time.Duration(nanosec100 / 10_000 * int64(time.Millisecond))
}
