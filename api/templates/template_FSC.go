// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDFSC = "tmpl/ui/fsc.tmpl"

//
// Renders HTML of template
// FSC with struct types.FSCs
func FSC(d types.FSCs) string {
	return netbFSC(d)
}

// Render template with JSON string as
// data.
func netFSC(args ...interface{}) string {

	// Get data from JSON
	var d = netcFSC(args...)
	return netbFSC(d)

}

// template render function
func netbFSC(d types.FSCs) string {
	localid := templateIDFSC
	name := "FSC"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcFSC(args ...interface{}) (d types.FSCs) {

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
