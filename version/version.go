package version

import (
	"fmt"
	"time"
)

// TimeFormat define time format
const TimeFormat = "Mon Jan 2 15:04:05 +0800 MST 2006"

var (
	// Version program version
	Version = "v1.0.0"

	// GitSHA Git SHA Value will be set during build
	GitSHA = "Not provided"

	// ReleaseDate latest git commit date
	ReleaseDate = time.Now()
)

// LogVersion returns formated version info
func LogVersion() string {
	return fmt.Sprintf("web version: %s, release date: %v, revision: %s", Version, ReleaseDate, GitSHA)
}
