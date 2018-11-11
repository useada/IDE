// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDMV = "tmpl/ui/mv.tmpl"

//
// Renders HTML of template
// MV with struct types.FSCs
func MV(d types.FSCs) string {
	return netbMV(d)
}

// Render template with JSON string as
// data.
func netMV(args ...interface{}) string {

	// Get data from JSON
	var d = netcMV(args...)
	return netbMV(d)

}

// template render function
func netbMV(d types.FSCs) string {
	localid := templateIDMV
	name := "MV"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcMV(args ...interface{}) (d types.FSCs) {

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
