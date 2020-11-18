package main

import (
	"fmt"
	"sort"
)

type Meeting struct {
	start, end int
}

type ByMeetingStart []*Meeting

func (ml ByMeetingStart) Less(i, j int) bool {
	return ml[i].start < ml[j].start
}

func (ml ByMeetingStart) Swap(i, j int) {
	ml[i], ml[j] = ml[j], ml[i]
}

func (ml ByMeetingStart) Len() int {
	return len(ml)
}

func (ml ByMeetingStart) Print() {
	for _, m := range ml {
		fmt.Printf("%v ", *m)
	}
	fmt.Println()
}

// {2, 3}, {6, 8}, {-5, -4}, {4, 5},
// {-5, -4} {2, 3} {4, 5} {6 8}
// {-8, 7, 1} => {-8, -5} {-4, 2} {3, 4} {5, 6}
func (ml ByMeetingStart) FindSlots(start, end int, duration int) []Meeting {
	freeSlots := []Meeting{}
	ss := start
	for _, m := range ml {
		if end <= m.start {
			break
		}
		if ss + duration <= m.start {
			slot := Meeting{ss, m.start}
			freeSlots = append(freeSlots, slot)
		}
		ss = m.end
	}
	return freeSlots	
}

func main() {
	meetings := []*Meeting{
		{2, 3}, {6, 8}, {-5, -4}, {4, 5},
	}
	ml := ByMeetingStart(meetings)
	ml.Print()
	sort.Sort(ml)
	ml.Print()
	freeSlots := ml.FindSlots(-8, 7, 1)
	fmt.Println("freeSlots => ", freeSlots)
}

