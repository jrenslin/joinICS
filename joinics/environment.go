package joinICS

import (
	jbasefuncs "github.com/jrenslin/jbasefuncs"
)

// Function to always be run when joinICS is run. Sets up joinICS's folder in the user's .config directory.
// Creates one file and  there:
// - config.json is the config file. Here, the input folders and the output folders are set.
func ensure_working_environment() {
	jbasefuncs.EnsureDir(baseDir)
	jbasefuncs.EnsureJson(baseLocation)
}
