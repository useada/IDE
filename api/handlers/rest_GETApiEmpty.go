package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	templates "github.com/thestrukture/IDE/api/templates"

	methods "github.com/thestrukture/IDE/api/methods"

	types "github.com/thestrukture/IDE/types"
)

func GETApiEmpty(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	methods.ClearLogs(r.FormValue("pkg"))
	response = templates.Alert(types.Alertbs{Type: "success", Text: "Your build logs are cleared."})

	callmet = true
	return
}