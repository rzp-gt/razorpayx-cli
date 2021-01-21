package fixtures

import (
	"fmt"
	"sort"

	"github.com/spf13/afero"
)

// Events is a mapping of pre-built trigger events and the corresponding json file
var Events = map[string]string{
	"payout.created": "/payout.created.json",
}

// BuildFromFixture creates a new fixture struct for a file
func BuildFromFixture(fs afero.Fs, apiKey, apiSecret, apiBaseURL, jsonFile string) (*Fixture, error) {
	fixture, err := NewFixture(
		fs,
		apiKey,
		apiSecret,
		apiBaseURL,
		jsonFile,
	)
	if err != nil {
		return nil, err
	}

	return fixture, nil
}

// EventList prints out a padded list of supported trigger events for printing the help file
func EventList() string {
	var eventList string
	for _, event := range EventNames() {
		eventList += fmt.Sprintf("  %s\n", event)
	}

	return eventList
}

// EventNames returns an array of all the event names
func EventNames() []string {
	names := []string{}
	for name := range Events {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

func reverseMap() map[string]string {
	reversed := make(map[string]string)
	for name, file := range Events {
		reversed[file] = name
	}

	return reversed
}
