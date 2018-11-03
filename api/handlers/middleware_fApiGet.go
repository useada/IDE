package handlers

import (
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheikhshift/gos/core"
	"github.com/gorilla/sessions"
	templates "github.com/thestrukture/IDE/api/templates"
	"gopkg.in/mgo.v2/bson"

	methods "github.com/thestrukture/IDE/api/methods"

	"github.com/thestrukture/IDE/api/globals"
	types "github.com/thestrukture/IDE/types"
)

func fApiGet(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	me := types.SoftUser{Email: "Strukture user", Username: "Strukture user"}
	if r.FormValue("type") == "0" {

		mpk := []bson.M{}

		apps := methods.GetApps()

		for _, v := range apps {
			if v.Name != "" {

				appCo := []types.PkgItem{}
				Childtm := []types.PkgItem{}

				var folders []types.PkgItem
				var pkgpath = core.TrimSuffix(os.ExpandEnv("$GOPATH"), "/") + "/src/" + v.Name + "/"
				if globals.Windows {
					pkgpath = strings.Replace(pkgpath, "/", "\\", -1)
				}

				if _, errr := os.Stat(pkgpath + "gos.gxml"); !os.IsNotExist(errr) {
					gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + v.Name + "/gos.gxml")
					for _, b := range v.Groups {
						tmpls := []types.PkgItem{}

						for _, tm := range gos.Templates.Templates {
							if tm.Bundle == b {
								tmpls = append(tmpls, types.PkgItem{Type: "5", AppID: v.Name, Icon: "fa fa-page", DType: "5&tmpl=" + b + "/" + tm.Name, Text: tm.Name, ID: v.Name + "@pkg:" + b + "/" + tm.Name})
							}
						}

						Childtm = append(Childtm, types.PkgItem{AppID: v.Name, Text: b, Icon: "fa fa-square", CType: "4&bundle=" + b, DType: "4&bundle=" + b, RType: "4&bundle=" + b, Children: tmpls})

					}

					_ = filepath.Walk(pkgpath+"web", func(path string, file os.FileInfo, _ error) error {
						//fmt.Println(path)
						if file.IsDir() {
							lpathj := strings.Replace(path, pkgpath+"web", "", -1)

							loca := types.PkgItem{AppID: v.Name, Text: lpathj, Icon: "fa fa-folder", Children: []types.PkgItem{}}

							loca.CType = "5&path=" + lpathj
							loca.DType = "6&isDir=Yes&path=" + lpathj

							loca.MType = "6&path=" + lpathj

							files, _ := ioutil.ReadDir(path)

							for _, f := range files {
								if !f.IsDir() {
									var mjk string
									mjk = strings.Replace(path, pkgpath+"web", "", -1) + "/" + f.Name()
									if globals.Windows {
										mjk = strings.Replace(mjk, "/", "\\", -1)
									}

									loca.Children = append(loca.Children, types.PkgItem{AppID: v.Name, Text: f.Name(), Icon: "fa fa-page", Type: "6", ID: v.Name + "@pkg:" + mjk, MType: "6&path=" + mjk, DType: "6&isDir=No&path=" + mjk})

								}
							}

							folders = append(folders, loca)

						}
						//fmt.Println(file,path,file.Name,file.IsDir())
						//   var loca PkgItem = types.PkgItem{AppID:v.Name,Text: file.Name(),Icon: "fa fa-folder"}

						return nil
					})

					appCo = append(appCo, types.PkgItem{AppID: v.Name, Text: "Template bundles", Icon: "fa fa-pencil-square", CType: "3", Children: Childtm})
					appCo = append(appCo, types.PkgItem{AppID: v.Name, Text: "Web Resources", CType: "5&path=/", Children: folders, Icon: "fa fa-folder"})

					appCo = append(appCo, types.PkgItem{AppID: v.Name, Type: "18", Text: "Testing", Icon: "fa fa-flask"})
					appCo = append(appCo, types.PkgItem{AppID: v.Name, Type: "8", Text: "Structs", Icon: "fa fa-share-alt"})
					//appCo = append(appCo, types.PkgItem{AppID:v.Name,Type:"9",Text: "Interface funcs",Icon: "fa fa-share-alt-square"} )
					appCo = append(appCo, types.PkgItem{Type: "10", AppID: v.Name, Text: "Template pipelines", Icon: "fa fa-exchange"})

					appCo = append(appCo, types.PkgItem{AppID: v.Name, Type: "11", Text: "Web services", Icon: "fa fa-circle-o-notch"})
				}

				var goFiles, ymlFiles []types.PkgItem

				_ = filepath.Walk(pkgpath, func(path string, file os.FileInfo, _ error) error {
					//fmt.Println(path)
					if file.IsDir() {
						lpathj := strings.Replace(path, pkgpath, "", -1)

						loca := types.PkgItem{AppID: v.Name, Text: lpathj, Icon: "fa fa-circle", Children: []types.PkgItem{}}
						hasgo := false
						files, _ := ioutil.ReadDir(path)
						for _, f := range files {
							if !f.IsDir() && strings.Contains(f.Name(), ".go") {

								var mjk string
								mjk = strings.Replace(path, pkgpath, "", -1) + "/" + f.Name()
								if globals.Windows {
									mjk = strings.Replace(mjk, "/", "\\", -1)
								}
								hasgo = true
								loca.Children = append(loca.Children, types.PkgItem{AppID: v.Name, Text: f.Name(), Icon: "fa fa-code", Type: "60", ID: v.Name + "@pkg:" + mjk, MType: "60&path=" + mjk, DType: "60&isDir=No&path=" + mjk})

							}
						}

						loca.CType = "50&path=" + lpathj
						loca.DType = "60&isDir=Yes&path=" + lpathj

						loca.MType = "60&path=" + lpathj

						if hasgo {
							goFiles = append(goFiles, loca)
						}

					}
					//fmt.Println(file,path,file.Name,file.IsDir())
					//   var loca PkgItem = types.PkgItem{AppID:v.Name,Text: file.Name(),Icon: "fa fa-folder"}

					return nil
				})

				_ = filepath.Walk(pkgpath, func(path string, file os.FileInfo, _ error) error {
					//fmt.Println(path)
					if file.IsDir() {
						lpathj := strings.Replace(path, pkgpath, "", -1)

						loca := types.PkgItem{AppID: v.Name, Text: lpathj, Icon: "fa fa-folder", Children: []types.PkgItem{}}
						hasyml := false
						files, _ := ioutil.ReadDir(path)
						for _, f := range files {
							if !f.IsDir() && strings.Contains(f.Name(), ".yml") {

								var mjk string
								mjk = strings.Replace(path, pkgpath, "", -1) + "/" + f.Name()
								if globals.Windows {
									mjk = strings.Replace(mjk, "/", "\\", -1)
								}
								hasyml = true
								loca.Children = append(loca.Children, types.PkgItem{AppID: v.Name, Text: f.Name(), Icon: "fa fa-magic", Type: "61", ID: v.Name + "@pkg:" + mjk, MType: "60&path=" + mjk, DType: "60&isDir=No&path=" + mjk})

							}
						}

						loca.CType = "51&path=" + lpathj
						loca.DType = "60&isDir=Yes&path=" + lpathj

						loca.MType = "60&path=" + lpathj

						if hasyml {
							ymlFiles = append(ymlFiles, loca)
						}

					}

					return nil
				})

				appCo = append(appCo, types.PkgItem{AppID: v.Name, Text: "Go SRC", CType: "50&path=/", Children: goFiles, Icon: "fa fa-cube"})

				appCo = append(appCo, types.PkgItem{AppID: v.Name, Type: "300", Text: "KanBan board", Icon: "fa fa-briefcase"})

				if v.Type != "faas" {

					appCo = append(appCo, types.PkgItem{AppID: v.Name, Type: "16", Text: "Logs", Icon: "fa fa-list"})
					appCo = append(appCo, types.PkgItem{AppID: v.Name, Type: "7", Text: "Build center", Icon: "fa fa-server"})
					appCo = append(appCo, types.PkgItem{AppID: v.Name, Type: "5500", Text: "Docker", Icon: "fa fa-cloud-upload"})

				} else {
					var functions []types.PkgItem

					for _, fn := range v.Groups {
						ref := types.PkgItem{AppID: v.Name, Text: fn, Icon: "fa fa-cube", Type: "62", ID: v.Name + "@pkg:" + fn, DType: "62&path=" + fn}
						functions = append(functions, ref)
					}

					appCo = append(appCo, types.PkgItem{AppID: v.Name, CType: "52", Children: functions, Text: "Functions", Icon: "fa fa-cubes"})
				}

				appCo = append(appCo, types.PkgItem{AppID: v.Name, CType: "51&path=/", Children: ymlFiles, Text: "YAML files", Icon: "fa fa-folder"})

				//appCo = append(appCo, types.PkgItem{AppID:v.Name,Type:"12",Text: "Timers",Icon: "fa fa-clock-o"} )

				rootel := bson.M{"dtype": "3", "text": v.Name, "type": "1", "id": v.Name, "children": appCo, "appid": v.Name, "btype": "on"}

				if v.Type == "webapp" {
					rootel["icon"] = "fa fa-globe"
				} else if v.Type == "app" {
					rootel["icon"] = "fa fa-folder"
					rootel["project"] = true
				} else if v.Type == "faas" {
					rootel["icon"] = "fa fa-rocket"
					rootel["project"] = true
					rootel["btype"] = nil
				} else {
					rootel["icon"] = "fa fa-gift"
				}

				//append to children
				//add server in
				mpk = append(mpk, rootel)
			}
		}

		response = mResponse(mpk)
	} else if r.FormValue("type") == "1" {

		//get package
		sapp := methods.GetApp(methods.GetApps(), r.FormValue("id"))
		prefix := "/api/put?type=0&id=" + sapp.Name

		//set params democss,port,key,name,type
		editor := types.SPackageEdit{Type: sapp.Type, TName: sapp.Name}
		pkgpath := os.ExpandEnv("$GOPATH") + "/src/" + sapp.Name + "/gos.gxml"

		if _, err := os.Stat(pkgpath); !os.IsNotExist(err) {

			gos, _ := core.LoadGos(pkgpath)
			editor.IType = types.Aput{Link: prefix, Param: "app", Value: gos.Type}

			editor.Port = types.Aput{Link: prefix, Param: "port", Value: gos.Port}
			editor.Key = types.Aput{Link: prefix, Param: "key", Value: gos.Key}
			editor.Domain = types.Aput{Link: prefix, Param: "domain", Value: gos.Domain}
			editor.Erpage = types.Aput{Link: prefix, Param: "erpage", Value: gos.ErrorPage}

			editor.Ffpage = types.Aput{Link: prefix, Param: "fpage", Value: gos.NPage}
			editor.Name = types.Aput{Link: prefix, Param: "Name", Value: sapp.Name}
			editor.Package = types.Aput{Link: "/api/put?type=16&pkg=" + sapp.Name, Param: "npk", Value: gos.Package}

			editor.Mainf = gos.Main
			editor.Shutdown = gos.Shutdown
			editor.Initf = gos.Init_Func
			editor.Sessionf = gos.Session

			varf := []types.Inputs{}
			varf = append(varf, types.Inputs{Name: "is", Type: "text", Text: "Variable type"})
			varf = append(varf, types.Inputs{Name: "name", Type: "text", Text: "Variable name"})
			editor.CreateVar = types.RPut{Count: "4", Link: "/api/create?type=0&pkg=" + sapp.Name, Inputs: varf, ListLink: "/api/get?type=2&pkg=" + sapp.Name}

			varf = []types.Inputs{}
			varf = append(varf, types.Inputs{Name: "src", Type: "text", Text: "Package path"})

			editor.CreateImport = types.RPut{Count: "6", Link: "/api/create?type=1&pkg=" + sapp.Name, Inputs: varf, ListLink: "/api/get?type=3&pkg=" + sapp.Name}
			varf = []types.Inputs{}
			varf = append(varf, types.Inputs{Name: "src", Type: "text", Text: "Path to css lib"})
			editor.Css = types.RPut{Count: "6", Link: "/api/create?type=2&pkg=" + sapp.Name, Inputs: varf, ListLink: "/api/get?type=4&pkg=" + sapp.Name}
			response = templates.PackageEdit(editor)
		} else {
			response = ""
		}

	} else if r.FormValue("type") == "2" {

		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		for _, v := range gos.Variables {
			varf := []types.Inputs{}
			varf = append(varf, types.Inputs{Name: "is", Type: "text", Text: "Variable type", Value: v.Type})
			varf = append(varf, types.Inputs{Name: "name", Type: "text", Text: "Variable name", Value: v.Name})
			response = response + templates.RPUT(types.RPut{DLink: "/api/delete?type=0&pkg=" + r.FormValue("pkg") + "&id=" + v.Name, Count: "4", Link: "/api/act?type=1&pkg=" + r.FormValue("pkg") + "&id=" + v.Name, Inputs: varf})
		}

	} else if r.FormValue("type") == "3" {

		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		for _, v := range gos.RootImports {
			varf := []types.Inputs{}
			varf = append(varf, types.Inputs{Name: "src", Type: "text", Text: "Package path", Value: v.Src})

			response = response + templates.RPUT(types.RPut{DLink: "/api/delete?type=1&pkg=" + r.FormValue("pkg") + "&id=" + v.Src, Count: "6", Link: "/api/act?type=2&pkg=" + r.FormValue("pkg") + "&id=" + v.Src, Inputs: varf})
		}

	} else if r.FormValue("type") == "4" {
		sapp := methods.GetApp(methods.GetApps(), r.FormValue("pkg"))

		for _, v := range sapp.Css {

			varf := []types.Inputs{}
			varf = append(varf, types.Inputs{Name: "src", Type: "text", Text: "Path to css lib", Value: v})

			response = response + templates.RPUT(types.RPut{DLink: "/api/delete?type=2&pkg=" + r.FormValue("pkg") + "&id=" + v, Count: "6", Link: "/api/act?type=3&pkg=" + r.FormValue("pkg") + "&id=" + v, Inputs: varf})
		}

	} else if r.FormValue("type") == "5" {
		id := strings.Split(r.FormValue("id"), "@pkg:")
		data, _ := ioutil.ReadFile(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/tmpl/" + id[1] + ".tmpl")

		data = []byte(html.EscapeString(string(data)))
		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

		template := methods.GetTemplate(gos.Templates.Templates, id[1])

		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Type: "text", Value: template.Struct, Name: "struct", Text: "Interface to use with template"})

		response = templates.TemplateEdit(types.TemplateEdits{SavesTo: "tmpl/" + id[1] + ".tmpl", ID: methods.RandTen(), PKG: r.FormValue("space"), Mime: "html", File: data, Settings: types.RPut{Link: "/api/put?type=2&id=" + id[1] + "&pkg=" + r.FormValue("space"), Inputs: varf, Count: "6"}})
	} else if r.FormValue("type") == "6" {
		id := strings.Split(r.FormValue("id"), "@pkg:")
		filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/web" + id[1]
		var ftype string
		if strings.Contains(filep, ".css") {
			ftype = "css"
		} else if strings.Contains(filep, ".js") {
			ftype = "javascript"
		} else if strings.Contains(filep, ".html") {
			ftype = "html"
		} else if strings.Contains(filep, ".tmpl") {
			ftype = "html"
			//add auto complete linking
		}
		data, _ := ioutil.ReadFile(filep)
		data = []byte(html.EscapeString(string(data)))
		response = templates.WebRootEdit(types.WebRootEdits{SavesTo: id[1], Type: ftype, File: data, ID: methods.RandTen(), PKG: r.FormValue("space")})

	} else if r.FormValue("type") == "60" {
		id := strings.Split(r.FormValue("id"), "@pkg:")
		filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/" + id[1]

		filep = strings.Replace(filep, "//", "/", -1)

		if globals.Windows {
			filep = strings.Replace(filep, "/", "\\", -1)
		}

		data, _ := ioutil.ReadFile(filep)
		data = []byte(html.EscapeString(string(data)))
		response = templates.WebRootEdittwo(types.WebRootEdits{SavesTo: id[1], Type: "golang", File: data, ID: methods.RandTen(), PKG: r.FormValue("space")})

	} else if r.FormValue("type") == "61" {
		id := strings.Split(r.FormValue("id"), "@pkg:")
		filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/" + id[1]

		filep = strings.Replace(filep, "//", "/", -1)

		if globals.Windows {
			filep = strings.Replace(filep, "/", "\\", -1)
		}

		data, _ := ioutil.ReadFile(filep)
		data = []byte(html.EscapeString(string(data)))
		response = templates.WebRootEdittwo(types.WebRootEdits{SavesTo: id[1], Type: "yaml", File: data, ID: methods.RandTen(), PKG: r.FormValue("space")})

	} else if r.FormValue("type") == "62" {
		id := strings.Split(r.FormValue("id"), "@pkg:")
		function := id[1] + "/handler.go"

		filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/" + function

		filep = strings.Replace(filep, "//", "/", -1)

		if globals.Windows {
			filep = strings.Replace(filep, "/", "\\", -1)
		}

		data, _ := ioutil.ReadFile(filep)
		data = []byte(html.EscapeString(string(data)))
		response = templates.WebRootEdittwo(types.WebRootEdits{SavesTo: function, Faas: true, Type: "golang", File: data, ID: methods.RandTen(), PKG: r.FormValue("space"), PreviewLink: id[1]})

	} else if r.FormValue("type") == "7" {
		sapp := methods.GetApp(methods.GetApps(), r.FormValue("space"))
		response = templates.ROC(types.SROC{Name: r.FormValue("space"), Build: sapp.Passed, Time: sapp.LatestBuild, Pid: sapp.Pid})
	} else if r.FormValue("type") == "8" {

		filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/structs.dsl"

		b, e := ioutil.ReadFile(filep)
		if e != nil {

			b = []byte("<gos> \n \n </gos> ")
		}

		data := html.EscapeString(string(b[:len(b)]))

		b = []byte(data)

		response = templates.StructEditor(types.VHuf{Edata: b, PKG: r.FormValue("space")})

	} else if r.FormValue("type") == "9" {

		filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/objects.dsl"

		b, e := ioutil.ReadFile(filep)
		if e != nil {

			b = []byte("<gos> \n \n </gos> ")
		}

		data := html.EscapeString(string(b[:len(b)]))

		b = []byte(data)

		response = templates.ObjectEditor(types.VHuf{Edata: b, PKG: r.FormValue("space")})

	} else if r.FormValue("type") == "10" {

		filep := os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/methods.dsl"

		b, e := ioutil.ReadFile(filep)
		if e != nil {

			b = []byte("<gos> \n \n </gos> ")
		}

		data := html.EscapeString(string(b[:len(b)]))

		b = []byte(data)

		response = templates.MethodEditor(types.VHuf{Edata: b, PKG: r.FormValue("space")})

	} else if r.FormValue("type") == "11" {

		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Name: "path", Type: "text", Text: "Endpoint path"})
		kput := types.RPut{ListLink: "/api/get?type=13&space=" + r.FormValue("space"), Inputs: varf, Count: "6", Link: "/api/put?type=7&space=" + r.FormValue("space")}
		response = templates.EndpointEditor(types.TEditor{CreateForm: kput, PKG: r.FormValue("space")})

	} else if r.FormValue("type") == "12" {
		varf := []types.Inputs{}
		varf = append(varf, types.Inputs{Name: "name", Type: "text", Text: "Timer name"})
		kput := types.RPut{ListLink: "/api/get?type=14&space=" + r.FormValue("space"), Inputs: varf, Count: "6", Link: "/api/put?type=8&space=" + r.FormValue("space")}
		response = templates.TimerEditor(types.TEditor{CreateForm: kput, PKG: r.FormValue("space")})
	} else if r.FormValue("type") == "13" {

		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

		for _, v := range gos.Endpoints.Endpoints {

			varf := []types.Inputs{}
			varf = append(varf, types.Inputs{Name: "path", Type: "text", Text: "Endpoint path", Value: v.Path})
			//varf = append(varf, types.Inputs{Name:"method", Type:"text",Text:"Endpoint method",Value:v.Method})
			varf = append(varf, types.Inputs{Name: "typ", Type: "text", Text: "Request type : GET,POST,PUT,DELETE,f,star...", Value: v.Type})

			response = response + templates.RPUT(types.RPut{DLink: "/api/delete?type=7&pkg=" + r.FormValue("space") + "&path=" + v.Id, Link: "/api/put?type=9&id=" + v.Id + "&pkg=" + r.FormValue("space"), Count: "12", Inputs: varf}) + methods.Addjsstr

		}

	} else if r.FormValue("type") == "13r" {

		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("pkg") + "/gos.gxml")

		for _, v := range gos.Endpoints.Endpoints {

			if v.Id == r.FormValue("id") {
				id := methods.RandTen()
				response = templates.TemplateEditTwo(types.TemplateEdits{SavesTo: "gosforceasapi/" + r.FormValue("id") + "++()/", ID: id, PKG: r.FormValue("pkg"), Mime: "golang", File: []byte(v.Method)})
			}
		}

	} else if r.FormValue("type") == "14" {

		gos, _ := core.PLoadGos(os.ExpandEnv("$GOPATH") + "/src/" + r.FormValue("space") + "/gos.gxml")

		for _, v := range gos.Timers.Timers {

			varf := []types.Inputs{}
			varf = append(varf, types.Inputs{Name: "name", Type: "text", Text: "Timer name", Value: v.Name})
			varf = append(varf, types.Inputs{Name: "interval", Type: "number", Text: "Interval", Value: v.Interval})
			varf = append(varf, types.Inputs{Name: "unit", Type: "text", Text: "Timer refresh unit", Value: v.Unit})
			varf = append(varf, types.Inputs{Name: "method", Type: "text", Text: "Method to execute", Value: v.Method})
			response = response + templates.RPUT(types.RPut{DLink: "/api/delete?type=8&pkg=" + r.FormValue("space") + "&name=" + v.Name, Link: "/api/put?type=10&id=" + v.Name + "&pkg=" + r.FormValue("space"), Count: "2", Inputs: varf})

		}

	} else if r.FormValue("type") == "15" {

		tempx := templates.USettings(types.USettings{StripeID: me.StripeID, LastPaid: "Date", Email: me.Email})
		response = templates.Modal(types.SModal{Title: "Account settings", Body: tempx, Color: "orange"})
	} else if r.FormValue("type") == "16" {

		response = templates.Debugger(types.DebugObj{PKG: r.FormValue("space"), Username: ""})
	} else if r.FormValue("type") == "17" {

		var tDebugNode types.DebugObj

		if r.FormValue("id") == "Server" {
			tDebugNode = types.DebugObj{Time: "Server", Bugs: []types.DebugNode{}}
			gp := os.ExpandEnv("$GOPATH")
			os.Chdir(gp + "/src/" + r.FormValue("space"))
			//main.log
			rlog, err := ioutil.ReadFile("main.log")
			if err != nil {
				tDebugNode.RawLog = err.Error()
			} else {
				tDebugNode.RawLog = string(rlog)
			}
		} else {
			logs := methods.GetLogs(r.FormValue("space"))

			for _, logg := range logs {
				if logg.Time == r.FormValue("id") {
					tDebugNode = logg
				}
			}
		}

		response = templates.DebuggerNode(tDebugNode)

	} else if r.FormValue("type") == "18" {

		response = templates.EndpointTesting(types.Dex{Misc: r.FormValue("space")})

	} else if r.FormValue("type") == "300" {

		response = templates.KanBan(types.Dex{Misc: r.FormValue("space")})

	} else if r.FormValue("type") == "5500" {
		response = templates.Docker(types.Dex{Misc: r.FormValue("space")})
	}

	callmet = true

	return
}