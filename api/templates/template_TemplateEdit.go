// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDTemplateEdit = "tmpl/ui/user/panel/templateEditor.tmpl"

//
// Renders HTML of template
// TemplateEdit with struct types.TemplateEdits
func TemplateEdit(d types.TemplateEdits) string {
	return netbTemplateEdit(d)
}

// Render template with JSON string as
// data.
func netTemplateEdit(args ...interface{}) string {

	// Get data from JSON
	var d = netcTemplateEdit(args...)
	return netbTemplateEdit(d)

}

// template render function
func netbTemplateEdit(d types.TemplateEdits) string {
	localid := templateIDTemplateEdit
	name := "TemplateEdit"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcTemplateEdit(args ...interface{}) (d types.TemplateEdits) {

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
