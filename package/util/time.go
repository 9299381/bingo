package util

import (
	"time"
)

type DateTime struct {
	Start time.Time
	End   time.Time
}

func GetToday() *DateTime {
	year, month, day := time.Now().Date()
	dateTime := &DateTime{}
	start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	dateTime.Start = start
	dateTime.End = start.AddDate(0, 0, 1)
	return dateTime
}
func GetYesterday() *DateTime {
	year, month, day := time.Now().Date()
	dateTime := &DateTime{}
	start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	dateTime.Start = start.AddDate(0, 0, -1)
	dateTime.End = start.AddDate(0, 0, 0)
	return dateTime
}
func GetTomorrow() *DateTime {
	year, month, day := time.Now().Date()
	dateTime := &DateTime{}
	start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	dateTime.Start = start.AddDate(0, 0, 1)
	dateTime.End = start.AddDate(0, 0, 2)
	return dateTime
}

func GetMonth() *DateTime {
	year, month, _ := time.Now().Date()
	dateTime := &DateTime{}
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	dateTime.Start = start
	dateTime.End = start.AddDate(0, 1, 0)
	return dateTime
}
func GetLastMonth() *DateTime {
	year, month, _ := time.Now().Date()
	dateTime := &DateTime{}
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	dateTime.Start = start.AddDate(0, -1, 0)
	dateTime.End = start.AddDate(0, 0, 0)
	return dateTime
}

func GetYear() *DateTime {
	year, _, _ := time.Now().Date()
	dateTime := &DateTime{}
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	dateTime.Start = start
	dateTime.End = start.AddDate(1, 0, 0)
	return dateTime
}
