package joinICS

import (
	"fmt"
	jbasefuncs "github.com/jrenslin/jbasefuncs"
	"strings"
)

func joinFiles(folder string) {

	var splitPath = strings.Split(folder, "/")
	var folderName = splitPath[len(splitPath)-1:]

	var output = "BEGIN:VCALENDAR\nVERSION:2.0\nCALSCALE:GREGORIAN\nPRODID:-//CALENDARSERVER.ORG//NONSGML Version 1//EN\n"
	for _, file := range jbasefuncs.Scandir(folder) {

		var contents = jbasefuncs.FileGetContents(file)

		var toKeep = []string{
			"SUMMARY",
			"DESCRIPTION",
			"DTSTART",
			"DTEND",
			"DTSTAMP",
		}

		var fileOutput = "BEGIN:VEVENT\n"

		for _, line := range strings.Split(contents, "\n") {
			for _, identifier := range toKeep {
				if len(line) > len(identifier) && line[0:len(identifier)] == identifier {
					fileOutput += line + "\n"
				}
			}
		}

		fileOutput += "END:VEVENT\n"

		output += fileOutput

	}

	output += "END:VCALENDAR"

	jbasefuncs.EnsureDir(Settings.OutputFolder)

	// Write to file
	jbasefuncs.FilePutContents(Settings.OutputFolder+folderName[0]+".ics", output)
	fmt.Println(Settings.OutputFolder + folderName[0] + ".ics")

}
