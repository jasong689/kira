package main

import (
	"strconv"
	"testing"
)

var event1 = Event{
	Id:    1,
	Start: 2,
	End:   6,
}

var event2 = Event{
	Id:    2,
	Start: 1,
	End:   3,
}

var event3 = Event{
	Id:    3,
	Start: 10,
	End:   15,
}

var event4 = Event{
	Id:    4,
	Start: 21,
	End:   90,
}

var event5 = Event{
	Id:    5,
	Start: 9,
	End:   22,
}

type orderedEvents struct {
	Events      []Event
	SortByStart []Event
	SortByEnd   []Event
}

var testOrderedEvents = orderedEvents{
	Events: []Event{
		event1,
		event2,
		event3,
		event4,
		event5,
	},
	SortByEnd: []Event{
		event2,
		event1,
		event3,
		event5,
		event4,
	},
	SortByStart: []Event{
		event2,
		event1,
		event5,
		event3,
		event4,
	},
}

func TestCreateTimeString(t *testing.T) {
	time := "15:31"
	stringValue := CreateTimeString(time)

	if stringValue != "1970-01-01T15:31:00Z" {
		t.Error("Expected 1970-01-01T15:31, got", stringValue)
	}
}

func TestToUnixTs(t *testing.T) {
	date := "2019-05-01"
	time := "13:30"

	if dateTime, _ := ToUnixTs(date, time); dateTime != 1556717400 {
		t.Error("Expected 1556717400, got", dateTime)
	}

	invalidTime := "14:91"

	if _, success := ToUnixTs(date, invalidTime); success {
		t.Error("Expected error on invalid time, got success")
	}

	invalidDate := "2019-05-35"

	if _, success := ToUnixTs(invalidDate, time); success {
		t.Error("Expected error on invalid date, got success")
	}
}

func TestSplitEventString(t *testing.T) {
	eventString := "2005-09-10 12:50 13:50"
	eventDetails := SplitEventString(eventString)

	if len(eventDetails) != 3 {
		t.Error("Expected 3 elements for event details, got", len(eventDetails))
	}

	if eventDetails[0] != "2005-09-10" {
		t.Error("Expected event date 2005-09-10, got", eventDetails[0])
	}

	if eventDetails[1] != "12:50" {
		t.Error("Expected event start 12:50, got", eventDetails[1])
	}

	if eventDetails[2] != "13:50" {
		t.Error("Expected event end 13:50, got", eventDetails[2])
	}
}

func TestCreateEvent(t *testing.T) {
	event, success := CreateEvent("2018-11-12 07:20 09:20", 5)

	if !success {
		t.Error("Expected success true, got false")
	}

	if event.Id != 5 {
		t.Error("Expected Id 5, got", event.Id)
	}

	if event.Start != 1542007200 {
		t.Error("Expected Start 1542007200, got", event.Start)
	}

	if event.End != 1542014400 {
		t.Error("Expected End 1542014400, got", event.End)
	}

	_, successShouldBeFalse := CreateEvent("2017-02-05 12:00 11:59", 0)

	if successShouldBeFalse {
		t.Error("Expected error, got succcess")
	}
}

func TestSortByStart(t *testing.T) {
	sortedEvents := SortByStart(testOrderedEvents.Events)

	for i, event := range sortedEvents {
		if event.Id != testOrderedEvents.SortByStart[i].Id {
			t.Error("Expected "+strconv.Itoa(testOrderedEvents.SortByStart[i].Id)+", got", event.Id)
		}
	}
}

func TestSortByEnd(t *testing.T) {
	sortedEvents := SortByEnd(testOrderedEvents.Events)

	for i, event := range sortedEvents {
		if event.Id != testOrderedEvents.SortByEnd[i].Id {
			t.Error("Expected "+strconv.Itoa(testOrderedEvents.SortByEnd[i].Id)+", got", event.Id)
		}
	}
}

func TestNearestEndIndexGreaterThanOrEqual(t *testing.T) {
	test1 := NearestEndIndexGreaterThanOrEqual(56, testOrderedEvents.SortByEnd)

	if test1 != 4 {
		t.Error("Expected 4, got", test1)
	}

	test2 := NearestEndIndexGreaterThanOrEqual(100, testOrderedEvents.SortByEnd)

	if test2 != -1 {
		t.Error("Expected -1, got", test2)
	}

	test3 := NearestEndIndexGreaterThanOrEqual(7, testOrderedEvents.SortByEnd)

	if test3 != 2 {
		t.Error("Expected 2, got", test3)
	}

	test4 := NearestEndIndexGreaterThanOrEqual(1, testOrderedEvents.SortByEnd)

	if test4 != 0 {
		t.Error("Expected 0, got", test4)
	}
}
