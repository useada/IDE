// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDNavbar = "tmpl/ui/navbar.tmpl"

//
// Renders HTML of template
// Navbar with struct types.Dex
func Navbar(d types.Dex) string {
	return netbNavbar(d)
}

// Render template with JSON string as
// data.
func netNavbar(args ...interface{}) string {

	// Get data from JSON
	var d = netcNavbar(args...)
	return netbNavbar(d)

}

// template render function
func netbNavbar(d types.Dex) string {
	localid := templateIDNavbar
	name := "Navbar"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcNavbar(args ...interface{}) (d types.Dex) {

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
