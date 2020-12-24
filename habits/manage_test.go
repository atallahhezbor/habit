package habits_test

import (
	"github.com/atallahhezbor/habit/habits"
	"github.com/spf13/viper"
	"os"
	"testing"
)

const testDataFile = ".test_data"

func init() {
	viper.Set("datafile", testDataFile)
}

func teardown() {
	pathToFile := os.Getenv("HOME") + "/" + testDataFile
	err := os.Remove(pathToFile)
	if err != nil {
		panic(err)
	}
}

func TestLoadEmpty(t *testing.T) {
	defer teardown()
	loaded := habits.Load()
	if len(loaded) != 0 {
		t.Error("Map initialized with non-zero length")
	}
}

func TestSave(t *testing.T) {
	defer teardown()
	loaded := habits.Load()
	loaded["test"] = habits.New("name", "shortname", "tag")
	habits.Save(loaded)
	reloaded := habits.Load()
	if len(reloaded) != 1 {
		t.Errorf("Map re-loaded with unexpected length %d, want %d", len(reloaded), 1)
	}
	if reloaded["test"].Name != "name" {
		t.Errorf("Map re-loaded with unexpected name %s, want %s", reloaded["test"].Name, "name")
	}
	if reloaded["test"].ShortName != "shortname" {
		t.Errorf("Map re-loaded with unexpected short name %s, want %s", reloaded["test"].ShortName, "shortname")
	}
	if reloaded["test"].Tag != "tag" {
		t.Errorf("Map re-loaded with unexpected tag %s, want %s", reloaded["test"].Tag, "tag")
	}
}
