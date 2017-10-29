// Function to check if the list of folders to be served is empty.
// If no, this means that the setup should/can be run.
package joinICS

import (
	"encoding/json"
	jbasefuncs "github.com/jrenslin/jbasefuncs"
)

type SettingsType struct {
	OutputFolder string   `json:"outputfolder"`
	InputFolders []string `json:"inputfolders"`
}

// Function for decoding the folder list.
func decodeSettings(filename string) SettingsType {
	file := jbasefuncs.FileGetContentsBytes(filename)

	var data SettingsType
	err := json.Unmarshal(file, &data)
	jbasefuncs.Check(err)

	return data
}

func storeSettings(inputFolders []string, outputFolder string) {

	Settings.InputFolders = inputFolders
	Settings.OutputFolder = outputFolder

	jbasefuncs.FilePutContents(baseLocation, ToJson(Settings))

}

func checkForSettings() bool {
	if len(Settings.OutputFolder) == 0 {
		return true
	} else {
		return false
	}

}
