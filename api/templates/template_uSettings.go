// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDuSettings = "tmpl/editor/settings.tmpl"

//
// Renders HTML of template
// uSettings with struct types.USettings
func USettings(d types.USettings) string {
	return netbuSettings(d)
}

// Render template with JSON string as
// data.
func netuSettings(args ...interface{}) string {

	// Get data from JSON
	var d = netcuSettings(args...)
	return netbuSettings(d)

}

// template render function
func netbuSettings(d types.USettings) string {
	localid := templateIDuSettings
	name := "uSettings"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcuSettings(args ...interface{}) (d types.USettings) {

	if len(args) > 0 {
		jsonData := args[0].(string)
		err := parseJSON(jsonData, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	}

	return
}
