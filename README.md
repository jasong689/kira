## Purpose
Given a sequence of events, each having a start and end time, this program will return the sequence of all pairs of overlapping events.

## Usage

Running the main.go will accept a list of "events" from stdin.

Events should be in the format of `2019-01-01 15:30 16:30`. Each event should be entered on a newline. Once all events have been entered and `end` has been sent to stdin, overlaps will be returned to stdout.


## Examples
Example 1

Input:
```
2018-12-01 11:50 12:50
2018-12-01 12:50 13:30
2018-12-05 14:30 16:40
2018-12-05 13:20 14:15
2018-12-01 13:20 14:20
```

Output:
```
(2018-12-01 12:50 13:30, 2018-12-01 13:20 14:20)
```


Example 2

Input:
```
2018-12-01 11:50 12:50
2018-12-01 12:50 13:30
2018-12-05 14:30 16:40
2018-12-05 13:20 14:15
2018-12-01 13:20 14:20
2018-12-06 15:16 19:21
2018-12-01 13:00 13:40
2018-12-05 14:10 15:01
```

Output:
```
(2018-12-01 12:50 13:30, 2018-12-01 13:00 13:40)
(2018-12-01 12:50 13:30, 2018-12-01 13:20 14:20)
(2018-12-01 13:00 13:40, 2018-12-01 13:20 14:20)
(2018-12-05 13:20 14:15, 2018-12-05 14:10 15:01)
(2018-12-05 14:10 15:01, 2018-12-05 14:30 16:40)
```