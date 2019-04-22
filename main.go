package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// Event - struct holding data relating to an event
type Event struct {
	Id    int
	Start int64
	End   int64
	Str   string
}

// Overlap - struct holding 2 events which overlap
type Overlap struct {
	Event1 Event
	Event2 Event
}

// ByStart - order by event start ascending
type ByStart []Event

// ByEnd - order by event end ascending
type ByEnd []Event

func (a ByStart) Len() int {
	return len(a)
}

func (a ByStart) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByStart) Less(i, j int) bool {
	return a[i].Start < a[j].Start
}

func (a ByEnd) Len() int {
	return len(a)
}

func (a ByEnd) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByEnd) Less(i, j int) bool {
	return a[i].End < a[j].End
}

// CreateTimeString - starting with unix epoch allows us to get seconds
// since beginning of the day
func CreateTimeString(timeString string) string {
	return "1970-01-01T" + timeString + ":00Z"
}

// ToUnixTs - convert date and time string to unix timestamp
func ToUnixTs(dateString string, timeString string) (int64, bool) {
	dateFormat := "2006-01-02"
	timeFormat := time.RFC3339

	dateParsed, dateErr := time.Parse(dateFormat, dateString)
	timeParsed, timeErr := time.Parse(timeFormat, CreateTimeString(timeString))

	if dateErr != nil || timeErr != nil {
		return -1, false
	}

	combinedDateTime := dateParsed.Add(time.Second * time.Duration(timeParsed.Unix()))

	return combinedDateTime.Unix(), true
}

// SplitEventString - split string of format 'Date Start End' into slice
func SplitEventString(eventString string) []string {
	return strings.Split(eventString, " ")
}

// CreateEvent - create event struct from an event string and id
func CreateEvent(eventString string, id int) (Event, bool) {
	details := SplitEventString(eventString)

	if len(details) < 3 {
		return Event{}, false
	}

	start, startSuccess := ToUnixTs(details[0], details[1])
	end, endSuccess := ToUnixTs(details[0], details[2])

	if !startSuccess || !endSuccess || start > end {
		return Event{}, false
	}

	return Event{
		Start: start,
		End:   end,
		Id:    id,
		Str:   eventString,
	}, true
}

// NearestEndIndexGreaterThanOrEqual - use binary search to find nearest event with and end time greater than or equal to
// the value we want
func NearestEndIndexGreaterThanOrEqual(value int64, events []Event) int {
	right := len(events) - 1
	left := 0

	if events[right].End < value {
		return -1
	}

	for left <= right {
		mid := left + ((right - left) / 2)
		val := events[mid].End

		if val == value {
			return mid
		} else if val > value {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left > right {
		return left
	}

	return right
}

// GetOverlaps - provided a list of events sorted by start time and end time
// returns a list of overlapping events
func GetOverlaps(eventsByStart []Event, eventsByEnd []Event) []Overlap {
	numOfEvents := len(eventsByStart)
	eventsByEndCopy := make([]Event, numOfEvents)
	copy(eventsByEndCopy, eventsByEnd)
	var overlaps []Overlap

	for i := 0; i < numOfEvents; i++ {
		currentEvent := eventsByStart[i]
		indexOfEventsWhereEndGreaterThanStart := NearestEndIndexGreaterThanOrEqual(currentEvent.Start, eventsByEndCopy)

		if indexOfEventsWhereEndGreaterThanStart < 0 {
			continue
		}

		checkEvents := eventsByEndCopy[indexOfEventsWhereEndGreaterThanStart:]
		for _, checkEvent := range checkEvents {
			if checkEvent.Start < currentEvent.End && checkEvent.Id != currentEvent.Id && checkEvent.Start > currentEvent.Start {
				overlaps = append(overlaps, Overlap{
					Event1: currentEvent,
					Event2: checkEvent,
				})
			}
		}
		// next search will have a start time greater than the current
		// so we only need to search any with end time later than our
		// current start time
		eventsByEndCopy = checkEvents
	}

	return overlaps
}

// SortByStart - sort events by start time ascending
func SortByStart(events []Event) []Event {
	sortedEvents := make([]Event, len(events))
	copy(sortedEvents, events)
	sort.Sort(ByStart(sortedEvents))

	return sortedEvents
}

// SortByEnd - sort events by end time ascending
func SortByEnd(events []Event) []Event {
	sortedEvents := make([]Event, len(events))
	copy(sortedEvents, events)
	sort.Sort(ByEnd(sortedEvents))

	return sortedEvents
}

func main() {
	fmt.Println("Enter events:")

	reader := bufio.NewReader(os.Stdin)
	var events []Event
	id := 0

	for {
		text, err := reader.ReadString('\n')
		eventString := strings.TrimSpace(text)

		if err == nil && eventString != "end" {
			if event, success := CreateEvent(eventString, id); success {
				events = append(events, event)
			} else {
				fmt.Println("Skipping invalid event: " + eventString)
			}
		} else {
			break
		}

		id++
	}

	sortedByStart := SortByStart(events)
	sortedByEnd := SortByEnd(events)
	overlaps := GetOverlaps(sortedByStart, sortedByEnd)

	if overlapLen := len(overlaps); overlapLen == 0 {
		fmt.Println("No overlaps found")
	}

	for _, overlap := range overlaps {
		fmt.Println("(" + overlap.Event1.Str + ", " + overlap.Event2.Str + ")")
	}
}
