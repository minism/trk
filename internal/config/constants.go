package config

import "time"

var UserTimeZone *time.Location

func init() {
	// TODO: This should be loaded from trk config or environment variable.
	var err error
	UserTimeZone, err = time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
}
