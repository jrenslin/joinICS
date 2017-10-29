// General json-related functions
package joinICS

import (
	"encoding/json"
	jbasefuncs "github.com/jrenslin/jbasefuncs"
)

func ToJson(p interface{}) string {
	bytes, err := json.MarshalIndent(p, "", "    ")
	jbasefuncs.Check(err)
	return string(bytes)
}
