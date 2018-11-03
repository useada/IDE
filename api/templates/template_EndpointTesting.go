package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"log"

	"github.com/fatih/color"
	"github.com/thestrukture/IDE/api/assets"
	"github.com/thestrukture/IDE/types"
)

// Render HTML of template
// EndpointTesting with struct types.Dex
func EndpointTesting(d types.Dex) string {
	return NetbEndpointTesting(d)
}

// recovery function used to log a
// panic.
func templateFNEndpointTesting(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (ui/endpointtester) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDEndpointTesting = "tmpl/ui/endpointtester.tmpl"

func NetEndpointTesting(args ...interface{}) string {

	localid := templateIDEndpointTesting
	var d *types.Dex
	defer templateFNEndpointTesting(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.Dex{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("EndpointTesting")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}
func bEndpointTesting(d types.Dex) string {
	return NetbEndpointTesting(d)
}

//

func NetbEndpointTesting(d types.Dex) string {
	localid := templateIDEndpointTesting
	defer templateFNEndpointTesting(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("EndpointTesting")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = types.Dex{}
	output.Reset()
	output = nil
	return outpescaped
}
func NetcEndpointTesting(args ...interface{}) (d types.Dex) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.Dex{}
	}
	return
}

func cEndpointTesting(args ...interface{}) (d types.Dex) {
	if len(args) > 0 {
		d = NetcEndpointTesting(args[0])
	} else {
		d = NetcEndpointTesting()
	}
	return
}