package validators

import (
	"errors"
)

// ArgValidator is an argument validator. It accepts a string and returns an
// error if the string is invalid, or nil otherwise.
type ArgValidator func(string) error

var (
	// ErrAPIKeyNotConfigured is the error returned when the loaded profile is missing the api key property
	ErrAPIKeyNotConfigured = errors.New("you have not configured API keys yet")
	// ErrDeviceNameNotConfigured is the error returned when the loaded profile is missing the device name property
	ErrDeviceNameNotConfigured = errors.New("you have not configured your device name yet")
)
