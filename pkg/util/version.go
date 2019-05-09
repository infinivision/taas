package util

import (
	"fmt"
)

// set on build time
var (
	GitCommit = ""
	BuildTime = ""
	GoVersion = ""
	Version   = ""
)

// PrintVersion print version info
func PrintVersion() bool {
	fmt.Println("Version  : ", Version)
	fmt.Println("GitCommit: ", GitCommit)
	fmt.Println("BuildTime: ", BuildTime)
	fmt.Println("GoVersion: ", GoVersion)
	return true
}
