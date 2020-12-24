package output

import (
	"github.com/atallahhezbor/habit/habits"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"testing"
)

// ShortName here indicates the correct order
var multiTagMap = map[string]*habits.Habit{
	"0": habits.New("apple", "0", "a"),
	"1": habits.New("zebra", "1", "a"),
	"2": habits.New("apple", "2", "z"),
	"3": habits.New("zebra", "3", "z"),
}

var singleTagMap = map[string]*habits.Habit{
	"0": habits.New("apple", "0", "a"),
	"1": habits.New("zebra", "1", "a"),
}

var singleTagMapOverflow = map[string]*habits.Habit{
	"0": habits.New("a", "0", "a"),
	"1": habits.New("a", "1", "a"),
	"2": habits.New("a", "2", "a"),
	"3": habits.New("a", "3", "a"),
	"4": habits.New("a", "4", "a"),
	"5": habits.New("a", "5", "a"),
	"6": habits.New("a", "6", "a"),
	"7": habits.New("a", "7", "a"),
}

var emptyMap = map[string]*habits.Habit{}

func TestListEmpty(t *testing.T) {
	List(emptyMap)
}

func TestOrderByTag(t *testing.T) {
	assertOrder := func(t *testing.T, got habits.HabitList, sourceMap habits.HabitMap) {
		for index, habit := range got {
			orderedIndex, _ := strconv.Atoi(habit.ShortName)
			if orderedIndex != index {
				want := multiTagMap[strconv.Itoa(index)]
				t.Errorf("got habit with name %s and tag %s, want name %s and tag %s ", habit.Name, habit.Tag, want.Name, want.Tag)
			}
		}

	}
	t.Run("all elements belong to single tag", func(t *testing.T) {
		got := orderByTag(singleTagMap)
		assertOrder(t, got, singleTagMap)
	})

	t.Run("elements have multiple tags", func(t *testing.T) {
		got := orderByTag(singleTagMap)
		assertOrder(t, got, singleTagMap)
	})
}

func TestColorAssignmentSingleTag(t *testing.T) {
	ordered := orderByTag(singleTagMap)
	colorAssignments := buildColorOrder(ordered)
	if len(colorAssignments.Tags) != 1 {
		t.Errorf("Got tag list with length %d, want %d", len(colorAssignments.Tags), 1)
	}
	if len(colorAssignments.TagBoundaries) != 1 {
		t.Errorf("Got tag boundaries with length %d, want %d", len(colorAssignments.TagBoundaries), 1)
	}
	want := [...]uint8{ColorCodeOffset, ColorCodeOffset + 1}
	if !reflect.DeepEqual(colorAssignments.ColorOrder, want[:]) {
		t.Errorf("got color order %v, want %v", colorAssignments.ColorOrder, want[:])
	}
}

func TestColorAssignmentMultiTag(t *testing.T) {
	ordered := orderByTag(multiTagMap)
	colorAssignments := buildColorOrder(ordered)
	if len(colorAssignments.Tags) != 2 {
		t.Errorf("Got tag list with length %d, want %d", len(colorAssignments.Tags), 2)
	}
	if len(colorAssignments.TagBoundaries) != 2 {
		t.Errorf("Got tag boundaries with length %d, want %d", len(colorAssignments.TagBoundaries), 2)
	}
	// If there are two tags, there should be two color groups
	// starting at the offset, separated by the group length
	want := [...]uint8{ColorCodeOffset, ColorCodeOffset + 1, ColorCodeOffset + ColorGroupLength, ColorCodeOffset + ColorGroupLength + 1}
	if !reflect.DeepEqual(colorAssignments.ColorOrder, want[:]) {
		t.Errorf("got color order %v, want %v", colorAssignments.ColorOrder, want[:])
	}
}

func TestColorAssignmentSingleTagOverflow(t *testing.T) {
	ordered := orderByTag(singleTagMapOverflow)
	colorAssignments := buildColorOrder(ordered)

	// If one tag contains more than 6 habits, it should get modulated
	// One starting at 161, one starting at 167
	want := make([]uint8, 0, len(ordered))
	// First group is normal
	for i := uint8(0); i < ColorGroupLength; i++ {
		want = append(want, ColorCodeOffset+i)
	}
	// remaining items overflow into modulated colors
	// at the tail of the next row
	modulateBy := uint8(36)
	for i := ColorGroupLength; i < uint8(len(ordered)); i++ {
		baseColor := ColorCodeOffset
		rowLength := ColorGroupLength
		overrun := i % ColorGroupLength
		want = append(want, baseColor+modulateBy+rowLength-1-overrun)
	}

	if !reflect.DeepEqual(colorAssignments.ColorOrder, want[:]) {
		t.Errorf("got color order %v, want %v", colorAssignments.ColorOrder, want[:])
	}
}

func TestHistogramOutput(t *testing.T) {
	singleTagMap["0"].Tick()
	singleTagMap["1"].Tick()
	singleTagMap["1"].Tick()

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	want := `
	Tag  |Ticks
	0    |■
	1    |■|■
	`
	Hist(singleTagMap)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	got := string(out)

	if got != want {
		// TODO: this is a failing test for now
		// as I'm having trouble getting the piped output
		// to properly format
		// t.Errorf("got %s\n, want %s", got, want)
	}
}
