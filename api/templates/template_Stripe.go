// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"github.com/thestrukture/IDE/types"
	"log"
)

// Template path
var templateIDStripe = "tmpl/ui/stripe.tmpl"

//
// Renders HTML of template
// Stripe with struct types.Dex
func Stripe(d types.Dex) string {
	return netbStripe(d)
}

// Render template with JSON string as
// data.
func netStripe(args ...interface{}) string {

	// Get data from JSON
	var d = netcStripe(args...)
	return netbStripe(d)

}

// template render function
func netbStripe(d types.Dex) string {
	localid := templateIDStripe
	name := "Stripe"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcStripe(args ...interface{}) (d types.Dex) {

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
