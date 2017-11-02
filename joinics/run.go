package joinICS

import (
	"fmt"
	"os"
	"strings"
)

func Run() {

	args := os.Args[1:]

	switch {
	case len(args) == 0 && checkForSettings() == false:

		for _, folder := range Settings.InputFolders {
			joinFiles(folder)
		}
	case len(args) == 3 && args[0] == "--set":
		inputFolders := strings.Split(args[1], "|")
		storeSettings(inputFolders, args[2])
		fmt.Println("Set settings.")
	default:
		fmt.Println("Missing settings")
		fmt.Println("Run joinICS --set <inputfolders> <outputfolder>,")
		fmt.Println("whereas the input folders must be split by a |")
	}

}
