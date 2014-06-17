package main

//
//import (
//	"time"
//)
//
//type Time struct {
//	// sec gives the number of seconds elapsed since
//	// January 1, year 1 00:00:00 UTC.
//	sec int64
//
//	// nsec specifies a non-negative nanosecond
//	// offset within the second named by Seconds.
//	// It must be in the range [0, 999999999].
//	//
//	// It is declared as uintptr instead of int32 or uint32
//	// to avoid garbage collector aliasing in the case where
//	// on a 64-bit system the int32 or uint32 field is written
//	// over the low half of a pointer, creating another pointer.
//	// TODO(rsc): When the garbage collector is completely
//	// precise, change back to int32.
//	nsec uintptr
//
//	// loc specifies the Location that should be used to
//	// determine the minute, hour, month, day, and year
//	// that correspond to this Time.
//	// Only the zero Time has a nil Location.
//	// In that case it is interpreted to mean UTC.
//	loc *Location
//}
//
//
//// A Month specifies a month of the year (January = 1, ...).
//type Month int
//
//const (
//	January Month = 1 + iota
//	February
//	March
//	April
//	May
//	June
//	July
//	August
//	September
//	October
//	November
//	December
//)
//
//var months = [...]string{
//	"January",
//	"February",
//	"March",
//	"April",
//	"May",
//	"June",
//	"July",
//	"August",
//	"September",
//	"October",
//	"November",
//	"December",
//}
//
//// String returns the English name of the month ("January", "February", ...).
//func (m Month) String() string { return months[m-1] }
//
//// Date returns the Time corresponding to
////	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
//// in the appropriate zone for that time in the given location.
////
//// The month, day, hour, min, sec, and nsec values may be outside
//// their usual ranges and will be normalized during the conversion.
//// For example, October 32 converts to November 1.
////
//// A daylight savings time transition skips or repeats times.
//// For example, in the United States, March 13, 2011 2:15am never occurred,
//// while November 6, 2011 1:15am occurred twice.  In such cases, the
//// choice of time zone, and therefore the time, is not well-defined.
//// Date returns a time that is correct in one of the two zones involved
//// in the transition, but it does not guarantee which.
////
//// Date panics if loc is nil.
//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
//	if loc == nil {
//		panic("time: missing time.Location in call to Date")
//	}
//
//	// Normalize month, overflowing into year.
//	m := int(month) - 1
//	year, m = norm(year, m, 12)
//	month = Month(m)+1
//
//	// Normalize nsec, sec, min, hour, overflowing into day.
//	sec, nsec = norm(sec, nsec, 1e9)
//	min, sec = norm(min, sec, 60)
//	hour, min = norm(hour, min, 60)
//	day, hour = norm(day, hour, 24)
//
//	y := uint64(int64(year) - absoluteZeroYear)
//
//	// Compute days since the absolute epoch.
//
//	// Add in days from 400-year cycles.
//	n := y / 400
//	y -= 400*n
//	d := daysPer400Years * n
//
//	// Add in 100-year cycles.
//	n = y/100
//	y -= 100*n
//	d += daysPer100Years*n
//
//	// Add in 4-year cycles.
//	n = y/4
//	y -= 4*n
//	d += daysPer4Years*n
//
//	// Add in non-leap years.
//	n = y
//	d += 365*n
//
//	// Add in days before this month.
//	d += uint64(daysBefore[month - 1])
//	if isLeap(year) && month >= March {
//		d++ // February 29
//	}
//
//	// Add in days before today.
//	d += uint64(day - 1)
//
//	// Add in time elapsed today.
//	abs := d * secondsPerDay
//	abs += uint64(hour * secondsPerHour + min * secondsPerMinute + sec)
//
//	unix := int64(abs) + (absoluteToInternal + internalToUnix)
//
//	// Look for zone offset for t, so we can adjust to UTC.
//	// The lookup function expects UTC, so we pass t in the
//	// hope that it will not be too close to a zone transition,
//	// and then adjust if it is.
////	_, offset, _, start, end := loc.lookup(unix)
////	if offset != 0 {
////		switch utc := unix - int64(offset); {
////		case utc < start:
////			_, offset, _, _, _ = loc.lookup(start-1)
////		case utc >= end:
////			_, offset, _, _, _ = loc.lookup(end)
////		}
////		unix -= int64(offset)
////	}
//
//	return Time{unix + unixToInternal, uintptr(nsec), loc}
//}
//
///**
//helper in built functions in Time.time
// */
//const (
//	// The unsigned zero year for internal calculations.
//	// Must be 1 mod 400, and times before it will not compute correctly,
//	// but otherwise can be changed at will.
//	absoluteZeroYear = -292277022399
//
//	// The year of the zero Time.
//	// Assumed by the unixToInternal computation below.
//	internalYear = 1
//
//	// The year of the zero Unix time.
//	unixYear = 1970
//
//	// Offsets to convert between internal and absolute or Unix times.
//	absoluteToInternal int64 = (absoluteZeroYear - internalYear) * 365.2425 * secondsPerDay
//	internalToAbsolute       = -absoluteToInternal
//
//	unixToInternal int64 = (1969 * 365 + 1969 / 4 - 1969 / 100 + 1969 / 400) * secondsPerDay
//	internalToUnix int64 = -unixToInternal
//)
//
//const (
//	secondsPerMinute = 60
//	secondsPerHour   = 60 * 60
//	secondsPerDay    = 24 * secondsPerHour
//	secondsPerWeek   = 7 * secondsPerDay
//	daysPer400Years  = 365 * 400 + 97
//	daysPer100Years  = 365 * 100 + 24
//	daysPer4Years    = 365 * 4 + 1
//)
//
//// norm returns nhi, nlo such that
////	hi * base + lo == nhi * base + nlo
////	0 <= nlo < base
//func norm(hi, lo, base int) (nhi, nlo int) {
//	if lo < 0 {
//		n := (-lo - 1) / base + 1
//		hi -= n
//		lo += n*base
//	}
//	if lo >= base {
//		n := lo / base
//		hi += n
//		lo -= n*base
//	}
//	return hi, lo
//}
//
//// daysBefore[m] counts the number of days in a non-leap year
//// before month m begins.  There is an entry for m=12, counting
//// the number of days before January of next year (365).
//var daysBefore = [...]int32{
//	0,
//	31,
//		31 + 28,
//			31 + 28 + 31,
//				31 + 28 + 31 + 30,
//					31 + 28 + 31 + 30 + 31,
//						31 + 28 + 31 + 30 + 31 + 30,
//							31 + 28 + 31 + 30 + 31 + 30 + 31,
//								31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
//									31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
//										31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
//											31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
//												31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
//}
//
//func isLeap(year int) bool {
//	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
//}
//
//// A Location maps time instants to the zone in use at that time.
//// Typically, the Location represents the collection of time offsets
//// in use in a geographical area, such as CEST and CET for central Europe.
//type Location struct {
//	name string
//	zone []zone
//	tx   []zoneTrans
//
//	// Most lookups will be for the current time.
//	// To avoid the binary search through tx, keep a
//	// static one-element cache that gives the correct
//	// zone for the time when the Location was created.
//	// if cacheStart <= t <= cacheEnd,
//	// lookup can return cacheZone.
//	// The units for cacheStart and cacheEnd are seconds
//	// since January 1, 1970 UTC, to match the argument
//	// to lookup.
//	cacheStart int64
//	cacheEnd   int64
//	cacheZone  *zone
//}
//
//// A zone represents a single time zone such as CEST or CET.
//type zone struct {
//	name   string // abbreviated name, "CET"
//	offset int    // seconds east of UTC
//	isDST  bool   // is this zone Daylight Savings Time?
//}
//
//// A zoneTrans represents a single time zone transition.
//type zoneTrans struct {
//	when         int64 // transition time, in seconds since 1970 GMT
//	index        uint8 // the index of the zone that goes into effect at that time
//	isstd, isutc bool  // ignored - no idea what these mean
//}
//
//
//
