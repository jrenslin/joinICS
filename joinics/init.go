package joinICS

import (
	"os/user"
)

func init() {

	user, _ := user.Current()
	baseDir = user.HomeDir + "/.config/joinICS"
	baseLocation = user.HomeDir + "/.config/joinICS/config.json"

	ensure_working_environment() // Ensure that config files exist

	Settings = decodeSettings(baseLocation)

}
