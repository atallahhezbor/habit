package habits

import (
	"time"
)

// Habit is a represenation of a repeated task and its metadata
// Maintains a slice of timestamp occurrences
type Habit struct {
	Name        string
	ShortName   string
	Tag         string
	Occurrences []time.Time
}

// New creates a new habit and return a pointer to it
func New(name, shortName, tag string) *Habit {
	return &Habit{name, shortName, tag, make([]time.Time, 0)}
}

// Tick indicates that an instance of the habit occurred by appending to `occurrences`
func (h *Habit) Tick() {
	h.Occurrences = append(h.Occurrences, time.Now())
}

// HabitList is an array of Habit pointers
type HabitList []*Habit

// HabitMap is a map of string to Habit pointers
type HabitMap map[string]*Habit

func (h HabitList) Len() int           { return len(h) }
func (h HabitList) Less(i, j int) bool { return h[i].Name < h[j].Name }
func (h HabitList) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
