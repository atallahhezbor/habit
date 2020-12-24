package habits_test

import (
	"github.com/atallahhezbor/habit/habits"
	"testing"
)

func TestCreate(t *testing.T) {
	name := "Test Habit"
	shortName := "th"
	tag := "test"
	h := habits.New(name, shortName, tag)
	if h.Name != name {
		t.Errorf("Created habit name mismatch. got %s want %s", h.Name, name)
	}
	if h.ShortName != shortName {
		t.Errorf("Created habit shortName mismatch. got %s want %s", h.ShortName, shortName)
	}
	if h.Tag != tag {
		t.Errorf("Created habit tag mismatch. got %s want %s", h.Tag, tag)
	}
	if len(h.Occurrences) != 0 {
		t.Error("Created habit had non-zero occurrence length")
	}
}

func TestTick(t *testing.T) {
	name := "Test Habit"
	shortName := "th"
	tag := "test"
	h := habits.New(name, shortName, tag)
	h.Tick()
	if len(h.Occurrences) != 1 {
		t.Errorf("Unexpected occurrence length after tick. got %d want %d", len(h.Occurrences), 1)
	}
	h.Tick()
	if len(h.Occurrences) != 2 {
		t.Errorf("Unexpected occurrence length after tick. got %d want %d", len(h.Occurrences), 2)
	}
}
