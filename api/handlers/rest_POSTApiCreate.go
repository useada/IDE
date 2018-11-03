package handlers

import (
	"net/http"
	"os"

	"github.com/cheikhshift/gos/core"
	"github.com/gorilla/sessions"
	templates "github.com/thestrukture/IDE/api/templates"

	methods "github.com/thestrukture/IDE/api/methods"

	types "github.com/thestrukture/IDE/types"
)

func POSTApiCreate(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	//me := &types.SoftUser{Email:"Strukture user", Username:"Strukture user"}

	if r.FormValue("type") == "0" {
		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		gos.Add("var", r.FormValue("is"), r.FormValue("name"))
		// fmt.Println(gos)

		gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

	} else if r.FormValue("type") == "1" {
		//import
		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		gos.Add("import", "", r.FormValue("src"))
		//fmt.Println(gos)

		gos.PSaveGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")
	} else if r.FormValue("type") == "2" {
		//css
		apps := methods.GetApps()
		app := methods.GetApp(methods.GetApps(), r.FormValue("pkg"))
		app.Css = append(app.Css, r.FormValue("src"))
		apps = methods.UpdateApp(methods.GetApps(), r.FormValue("pkg"), app)
		methods.SaveApps(apps)
		//Users.Update(bson.M{"uid":me.UID}, me)
	} else if r.FormValue("type") == "3" {
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Name: "name", Type: "text", Text: "Bundle name"})

		response = templates.Form(types.Forms{Link: "/api/act?type=4&pkg=" + r.FormValue("pkg"), CTA: "Create Bundle", Class: "warning", Inputs: varf})

	} else if r.FormValue("type") == "4" {
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Name: "name", Type: "text", Text: "Template name"})

		response = templates.Form(types.Forms{Link: "/api/act?type=5&pkg=" + r.FormValue("pkg") + "&bundle=" + r.FormValue("bundle"), CTA: "Create Template file", Class: "warning", Inputs: varf})

	} else if r.FormValue("type") == "5" {
		//prefix pkg
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Type: "text", Name: "path", Text: "Path"})
		varf = append(varf, types.Inputs{Type: "hidden", Name: "basesix"})
		varf = append(varf, types.Inputs{Type: "hidden", Name: "fmode", Value: "touch"})

		response = templates.FSC(types.FSCs{Path: r.FormValue("path"), Form: types.Forms{Link: "/api/act?type=6&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Create", Class: "warning"}})
	} else if r.FormValue("type") == "50" {
		//prefix pkg
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Type: "text", Name: "path", Text: "Path"})
		varf = append(varf, types.Inputs{Type: "hidden", Name: "basesix"})
		varf = append(varf, types.Inputs{Type: "hidden", Name: "fmode", Value: "touch"})

		response = templates.FSC(types.FSCs{Path: r.FormValue("path"), Form: types.Forms{Link: "/api/act?type=60&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Create", Class: "warning"}})
	} else if r.FormValue("type") == "51" {
		//prefix pkg
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Type: "text", Name: "path", Text: "Path"})
		varf = append(varf, types.Inputs{Type: "hidden", Name: "basesix"})
		varf = append(varf, types.Inputs{Type: "hidden", Name: "fmode", Value: "touch"})

		response = templates.FSC(types.FSCs{Path: r.FormValue("path"), Form: types.Forms{Link: "/api/act?type=61&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Add YAML file", Class: "warning"}})
	} else if r.FormValue("type") == "52" {
		//prefix pkg
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Type: "text", Name: "name", Text: "Function name."})
		response = templates.FSC(types.FSCs{Path: r.FormValue("path"), Hide: true, Form: types.Forms{Link: "/api/act?type=62&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Add function", Class: "success"}})

	} else if r.FormValue("type") == "6" {
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Type: "text", Name: "path", Misc: "required", Text: "New path"})

		response = templates.MV(types.FSCs{Path: r.FormValue("path"), Form: types.Forms{Link: "/api/act?type=7&pkg=" + r.FormValue("pkg") + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Move", Class: "warning"}})
	} else if r.FormValue("type") == "60" {
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Type: "text", Name: "path", Misc: "required", Text: "New path"})

		response = templates.MV(types.FSCs{Path: r.FormValue("path"), Form: types.Forms{Link: "/api/act?type=70&pkg=" + r.FormValue("pkg") + "&folder=" + "&prefix=" + r.FormValue("path"), Inputs: varf, CTA: "Move", Class: "warning"}})
	}

	callmet = true
	return
}