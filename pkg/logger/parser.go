package logger

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"regexp"
)

const (
	pattern string = "^[a-z0-9_-]*$"
	layout string = "15:04"
)

func ParseFile(filepath string) (*Club, error) {
	file, err := os.Open(filepath)
    if err != nil {
		return nil, fmt.Errorf("error opening file:  %s", err)
    }
    defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	club := &Club{
		Clients: make(map[string]bool),
		Queue: make([]string, 0, 1),
		Result: make([]Event, 0, 1),
	}

	// Read the number of tables
	if scanner.Scan() {
		line = scanner.Text()
		numTables, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid number of tables format: %s", line)
		}
		if numTables <= 0 {
			return nil, fmt.Errorf("invalid number of tables: %s", line)
		}
		club.Tables= make([]Table, numTables)
	}

	// Read the start and end times
	if scanner.Scan() {
		line = scanner.Text()
		times := strings.Split(line, " ")
		if len(times) != 2 {
			return nil, fmt.Errorf("invalid start and end times: %s", line)
		}
		club.StartTime, err = time.Parse(layout, times[0])
		if err != nil {
			return nil, fmt.Errorf("invalid time format: %s", line)
		}
		club.EndTime, err = time.Parse(layout, times[1])
		if err != nil {
			return nil, fmt.Errorf("invalid time format: %s", line)
		}
	}

	// Read the hourly rate
	if scanner.Scan() {
		line = scanner.Text()
		rate, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid hourly rate format: %s", line)
		}
		if rate <= 0 {
			return nil, fmt.Errorf("invalid hourly rate: %s", line)
		}
		club.HourlyRate = rate
	}

	// Read and parse the events
	for scanner.Scan() {
		if line = scanner.Text(); line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 3 || len(fields) > 4 {
			return nil, fmt.Errorf("invalid event format: %s ", line)
		}

		eventTime, err := time.Parse(layout, fields[0])
		if err != nil {
			return nil, fmt.Errorf("invalid event time format: %s ", line)
		}
		if len(club.Events) != 0 && !eventTime.After(club.Events[len(club.Events) - 1].Time){
			return nil,fmt.Errorf("invalid event time: %s ", line)
		}
		eventID, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("invalid event ID: %s ", line)
		}
		regex := regexp.MustCompile(pattern)
		if !regex.MatchString(fields[2]) {
			return nil, fmt.Errorf("invalid client name: %s ", line)
		}
		clientName := fields[2]
		table := 0
		if len(fields) == 4 {
			table, err = strconv.Atoi(fields[3])
			if err != nil {
				return nil, fmt.Errorf("invalid table num format: %s ", line)
			}
			if table < 0 || table > len(club.Tables) {
				return nil, fmt.Errorf("invalid table num: %s ", line)
			}
		}
		club.Events = append(club.Events, Event{
			Time: eventTime,
			ID:   eventID,
			Client:    clientName,
			TableNum: table,
		})
	}

	return club, nil
}

func (c *Club) PrintResults() {
	fmt.Println(c.StartTime.Format(layout))
	for _, event := range c.Result {
		if event.TableNum == 0 {
			fmt.Printf("%s %d %s\n", event.Time.Format(layout), event.ID, event.Client)
		} else {
			fmt.Printf("%s %d %s %d\n", event.Time.Format(layout), event.ID, event.Client, event.TableNum)
		}
	}
	fmt.Println(c.EndTime.Format(layout))
	for i, table := range c.Tables {
		fmt.Printf("%d %d %s\n", i+1, table.Income, fmtDuration(c.Tables[i].Duration))
	}
}