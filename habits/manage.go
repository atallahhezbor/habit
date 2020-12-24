package habits

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	"os"
)

/*
 TODO: Persist both a map for lookups AND an array for ordered display?
 ___
|map|  habits keyed by short name

 ___
|arr| list of habits ordered by name

 ___
 |map|  ordered habitLists keyed by tag
*/

var handle *codec.MsgpackHandle = new(codec.MsgpackHandle)
var sourcefile string
var path string

func init() {
	sourcefile = "/" + viper.GetString("datafile")
	path = os.Getenv("HOME") + sourcefile
}

// Load file if present, otherwise create and return empty map
func Load() (habitMap map[string]*Habit) {
	// TODO: way to make tests work without re-reading from config here
	sourcefile := "/" + viper.GetString("datafile")
	path := os.Getenv("HOME") + sourcefile

	habitMap = map[string]*Habit{}

	reader, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)

	defer func() {
		if err := reader.Close(); err != nil {
			panic(err)
		}
	}()
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	dec := codec.NewDecoderBytes(content, handle)
	dec.Decode(&habitMap)
	return habitMap
}

// Save persists the modified map to a file
func Save(habitMap map[string]*Habit) {
	// TODO: way to make tests work without re-reading from config here
	sourcefile := "/" + viper.GetString("datafile")
	path := os.Getenv("HOME") + sourcefile
	w, err := os.OpenFile(path, os.O_WRONLY, 0777) // TODO: be nicer here
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			panic(err)
		}
	}()
	enc := codec.NewEncoder(w, handle)
	err = enc.Encode(habitMap)
	fmt.Println(err)
}
