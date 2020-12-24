package output

import (
	"fmt"
	"github.com/atallahhezbor/habit/habits"
	"github.com/gookit/color"
	"github.com/ryanuber/columnize"
	"math"
	"sort"
	"strings"
	"time"
)

// ColorGroupLength defines groupings of ANSI colors
const ColorGroupLength uint8 = 6

// ColorCodeOffset defines where on the ANSI color chart to start
var ColorCodeOffset uint8 = 166

func init() {
	// TODO: accept color overrides from config, perhaps for light terminal themes
	// viper.SetDefault("color-overrides", DEFAULT_COLORS)
	// TODO: accept offset from config
}

// Turns a map of habits into a slice ordered by tag
func orderByTag(habitMap habits.HabitMap) habits.HabitList {
	habits := make(habits.HabitList, len(habitMap))
	index := 0
	for _, v := range habitMap {
		habits[index] = v
		index++
	}
	sort.Slice(habits, func(i, j int) bool {
		if habits[i].Tag < habits[j].Tag {
			return true
		} else if habits[i].Tag > habits[j].Tag {
			return false
		} else {
			return habits[i].Name < habits[j].Name
		}
	})
	return habits
}

// List colorizes and display habits by tag
func List(habitMap habits.HabitMap) {
	orderedHabits := orderByTag(habitMap)
	habitStrings := make([]string, len(habitMap), len(habitMap))
	// Maintain a map of tag to indexes that have that tag
	type HabitEntry struct {
		Index     int
		DaysSince int
	}
	byTag := map[string][]HabitEntry{}

	for index, h := range orderedHabits {
		numOccurrences := len(h.Occurrences)
		daysSince := -1
		if numOccurrences > 0 {
			lastOccurrence := h.Occurrences[numOccurrences-1]
			hoursSince := time.Now().Sub(lastOccurrence).Hours()
			daysSince = int(math.Floor(hoursSince)) / 24
		}
		// TODO: Optimization, persist this formatted string in Habit objects
		// But there's trickiness since the index is part of this string
		habitStrings[index] = fmt.Sprintf("%d. | %s | %s | ^%d days since last tick\n", index, h.Name, h.ShortName, daysSince)
		byTag[h.Tag] = append(byTag[h.Tag], HabitEntry{index, daysSince})
	}

	columns := columnize.SimpleFormat(habitStrings)
	lines := strings.Split(columns, "\n")

	colorAssignments := buildColorOrder(orderedHabits)

	habitIndex := 0
	for _, boundary := range colorAssignments.TagBoundaries {
		tag := orderedHabits[habitIndex].Tag
		fmt.Printf("%4v#", "")
		color.OpUnderscore.Println(tag)
		for habitIndex < boundary {
			line := lines[habitIndex]
			parts := strings.Split(line, "^")
			colorToUse := colorAssignments.ColorOrder[habitIndex]
			color.S256(colorToUse).Printf("%8v %s", "", parts[0])
			color.Gray.Print(parts[1])
			fmt.Println()
			habitIndex++
		}
	}
}

// ColorAssignments define groups of
type ColorAssignments struct {
	ColorOrder    []uint8
	TagBoundaries []int
	Tags          []string
}

// Given a slice of habits that is ordered by tag and name
// return a slice for each tag that contains a slice of color assignments
// for items of that tag
// TODO: could be nice to use reflection to color by
// an arbitrary field
func buildColorOrder(habits habits.HabitList) *ColorAssignments {
	numHabits := len(habits)

	colorOrder := make([]uint8, numHabits)
	boundaries := make([]int, 0, numHabits)
	tags := make([]string, 0, numHabits)
	ca := &ColorAssignments{colorOrder, boundaries, tags}

	if numHabits == 0 {
		return ca
	}

	// Build color pallete from base color
	// Each group of 6x3 is a pallete.
	// Offset 160 dark colors
	tagChanged := true
	modulate := false
	var tagIndex, habitIndex, modulation uint8 = 0, 0, 0
	for index, habit := range habits {
		// Reset groupings for each new tag
		if index > 0 && habit.Tag != habits[index-1].Tag {
			tagChanged = true
			tagIndex++
			habitIndex = 0
			ca.TagBoundaries = append(ca.TagBoundaries, index)
			ca.Tags = append(ca.Tags, habit.Tag)
			modulate = false
			modulation = 0
		}
		// If there are more than 6 habits per tag
		// "modulate" to the next row of ansi colors
		// and zig zag back towards the start of the group
		if modulate {
			modulation += 36 + ColorGroupLength - habitIndex%ColorGroupLength - 1
		}
		if tagChanged {
			baseColor := uint8(tagIndex*ColorGroupLength + ColorCodeOffset)
			colorAssignment := baseColor + uint8(habitIndex) + modulation
			ca.ColorOrder[index] = colorAssignment
		}
		habitIndex++
	}
	// Edge case for the last tag where there is no change
	ca.TagBoundaries = append(ca.TagBoundaries, numHabits)
	ca.Tags = append(ca.Tags, habits[numHabits-1].Tag)
	return ca
}

// Hist groups and colorizes habit occurences, optionally by category
func Hist(habitMap habits.HabitMap) {
	taskHeader := "Task"
	tickHeader := "Ticks"
	orderedHabits := orderByTag(habitMap)
	colorAssignments := buildColorOrder(orderedHabits)

	// Format padding of the headers
	maxKeyLength := len(taskHeader)
	for _, habit := range habitMap {
		shortNameLength := len(habit.ShortName)
		if shortNameLength > maxKeyLength {
			maxKeyLength = shortNameLength
		}
	}
	paddingStringFmt := fmt.Sprintf("%%-%dv|", maxKeyLength)
	color.OpBold.Printf(paddingStringFmt, taskHeader)
	color.OpBold.Print(tickHeader)
	fmt.Println()

	// TODO: be smart about aggregation and density
	// i.e. different symbol, more intense color, etc
	for habitIndex, colorToUse := range colorAssignments.ColorOrder {
		habit := orderedHabits[habitIndex]
		fmt.Printf(paddingStringFmt, habit.ShortName)
		for i := 0; i < len(habit.Occurrences); i++ {
			color.S256(colorToUse).Print("■")
			fmt.Print("|")
		}
		fmt.Println()
	}
}
