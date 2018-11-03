package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/sessions"
	methods "github.com/thestrukture/IDE/api/methods"
)

func POSTApiDockerfile(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	imageName := r.FormValue("image")
	strat := r.FormValue("strat")
	port := r.FormValue("port")
	pkg := r.FormValue("pkg")

	dockerFilePath := filepath.Join(os.ExpandEnv("$GOPATH"), "src", pkg, "Dockerfile")

	var dockerfile string

	if strat == "Fast" {
		dockerfile = fmt.Sprintf(methods.DockerLarge, imageName, port, port)
	} else {
		dockerfile = fmt.Sprintf(methods.DockerSmall, imageName, port, port)
	}

	err := ioutil.WriteFile(dockerFilePath, []byte(dockerfile), 0700)

	if err != nil {
		log.Println(err)
	}

	response = "OK"

	callmet = true
	return
}