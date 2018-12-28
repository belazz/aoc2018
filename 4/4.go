package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Schedule struct {
	 minutesSlept int // to figure out which has slept the most
	 minutes [60]int // to figure out overlapping minutes on different days
}

type Record struct {
	when time.Time
	action string
}

func main() {
	guards := make(map[int]*Schedule)
	records := []Record{}

	regex := regexp.MustCompile(`\[(\d+-\d+-\d+ \d+:\d+)\] (.+)`)
	idRegex := regexp.MustCompile(`(\d+)`)

	contents, err := ioutil.ReadFile("4/input.txt")
	if err != nil {
		fmt.Println("input file not found")
	}
	matches := regex.FindAllSubmatch(contents, -1)

	for _, fullmatch := range matches {
		dateStr := string(fullmatch[1])
		str := string(fullmatch[2])
		datetime, err := time.Parse("2006-01-02 15:04", dateStr)
		if err == nil {
			records = append(records, Record{datetime, str})
		} else {
			fmt.Println("cant parse given time")
			fmt.Println(err)
		}
	}


	sort.Slice(records, func(i, j int) bool {
		return records[i].when.Before(records[j].when)
	})
	currentGuard := -1
	startedSleeping := 0
	// filling guards with schedule -> id:Schedule map
	for _, record := range records {
		id, _ := strconv.Atoi(idRegex.FindString(record.action))
		if id != 0 {
			currentGuard = id
		} else {
			switch {
			case strings.Contains(record.action, "falls asleep"):
				startedSleeping = record.when.Minute()
				if _, ok := guards[currentGuard]; ok == false {
					guards[currentGuard] = &Schedule{0, [60]int{}}
				}
			case strings.Contains(record.action, "wakes up"):
				guards[currentGuard].minutesSlept += record.when.Minute() - startedSleeping
				for i := startedSleeping; i < record.when.Minute(); i++ {
					guards[currentGuard].minutes[i]++
				}
			}
		}
	}
	// part1
	maxMinutesSlept := 0
	maxGuardId := 0
	maxCount := 0
	whichMinute := 0

	// part2
	mostOftenSleptGuardId := 0
	mostOftenSleptMinute := 0

	for id, schedule := range guards {
		// calculating max slept duration and the guard id that has slept the most
		if schedule.minutesSlept > maxMinutesSlept {
			maxMinutesSlept = schedule.minutesSlept
			maxGuardId = id
		}
		// calculating most often slept minute and the guard that most often slept on particular minute
		for minute, count := range schedule.minutes {
			if count > maxCount {
				mostOftenSleptGuardId = id
				maxCount = count
				mostOftenSleptMinute = minute
			}
		}
	}

	// back to zero after getting answer to part2
	// again calculating most often slept minute BUT for the guard that slept the most
	maxCount = 0
	for minute, count := range guards[maxGuardId].minutes {
		if count > maxCount {
			maxCount = count
			whichMinute = minute
		}
	}

	fmt.Printf("Par1 answer: id: %v, minute: %v, result:  %v\n", maxGuardId, whichMinute, maxGuardId * whichMinute)
	fmt.Printf("Par2 answer: id: %v, minute: %v, result:  %v", mostOftenSleptGuardId, mostOftenSleptMinute, mostOftenSleptGuardId * mostOftenSleptMinute)
}



