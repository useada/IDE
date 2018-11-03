package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	gosweb "github.com/cheikhshift/gos/web"
	"github.com/gorilla/sessions"

	sessionStore "github.com/thestrukture/IDE/api/sessions"
)

var WebCache = gosweb.NewCache()

func mResponse(v interface{}) string {
	data, _ := json.Marshal(&v)
	return string(data)
}

func ApiAttempt(w http.ResponseWriter, r *http.Request) (callmet bool) {
	var response string
	var session *sessions.Session

	var er error

	if session, er = sessionStore.Store.Get(r, "session-"); er != nil {
		session, _ = sessionStore.Store.New(r, "session-")
	}

	if strings.Contains(r.URL.Path, "/api/get") {
		response, callmet = fApiGet(w, r, session)
	}
	if r.Method == "RESET" {
		return true
	} else if isURL := (r.URL.Path == "/api/dockerfile" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiDockerfile(w, r, session)

	} else if isURL := (r.URL.Path == "/api/composer" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiComposer(w, r, session)

	} else if !callmet && gosweb.UrlAtZ(r.URL.Path, "/api/socket") {

		response, callmet = ApiSocket(w, r, session)

	} else if isURL := (r.URL.Path == "/api/pkg-bugs" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		response, callmet = GETApiPkgBugs(w, r, session)

	} else if isURL := (r.URL.Path == "/api/kanban" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		response, callmet = GETApiKanban(w, r, session)

	} else if isURL := (r.URL.Path == "/api/git" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiGit(w, r, session)

	} else if isURL := (r.URL.Path == "/api/kanban" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiKanban(w, r, session)

	} else if isURL := (r.URL.Path == "/api/empty" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		response, callmet = GETApiEmpty(w, r, session)

	} else if isURL := (r.URL.Path == "/api/tester/" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiTester(w, r, session)

	} else if isURL := (r.URL.Path == "/api/create" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiCreate(w, r, session)

	} else if isURL := (r.URL.Path == "/api/delete" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiDelete(w, r, session)

	} else if isURL := (r.URL.Path == "/api/rename" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiRename(w, r, session)

	} else if isURL := (r.URL.Path == "/api/new" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiNew(w, r, session)

	} else if isURL := (r.URL.Path == "/api/act" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiAct(w, r, session)

	} else if isURL := (r.URL.Path == "/api/put" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiPut(w, r, session)

	} else if isURL := (r.URL.Path == "/api/build" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		response, callmet = GETApiBuild(w, r, session)

	} else if isURL := (r.URL.Path == "/api/start" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		response, callmet = GETApiStart(w, r, session)

	} else if isURL := (r.URL.Path == "/api/stop" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		response, callmet = GETApiStop(w, r, session)

	} else if isURL := (r.URL.Path == "/api/bin" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		response, callmet = GETApiBin(w, r, session)

	} else if isURL := (r.URL.Path == "/api/export" && r.Method == strings.ToUpper("GET")); !callmet && isURL {

		response, callmet = GETApiExport(w, r, session)

	} else if !callmet && gosweb.UrlAtZ(r.URL.Path, "/api/complete") {

		response, callmet = ApiComplete(w, r, session)

	} else if isURL := (r.URL.Path == "/api/console" && r.Method == strings.ToUpper("POST")); !callmet && isURL {

		response, callmet = POSTApiConsole(w, r, session)

	} else if !callmet && gosweb.UrlAtZ(r.URL.Path, "/api/terminal_realtime") {

		response, callmet = ApiTerminal_realtime(w, r, session)

	}
	if callmet {
		session.Save(r, w)
		session = nil
		if response != "" {
			//Unmarshal json
			//w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
		}
		return
	}
	session = nil
	return
}