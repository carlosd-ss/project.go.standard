// Go Api server
// @jeffotoni

package handler

import (
	"net/http"
	"strings"

	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/util"
)

//HomeHandler ..if len(ENVAMBI) <= 0 {
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		///////////////////////////////////
		// to create routes dynamically
		// with regex for querys use in endpoints
		endpoint, nameregex := util.EndpointRegex(r)
		if endpoint == "/user" && len(nameregex) >= 0 && http.MethodDelete == strings.ToUpper(r.Method) {
			//nameregex = crypt.Base64Dec(nameregex)
			UserDelete(nameregex, w, r)
			return
		}
		if endpoint == "/user" && len(nameregex) >= 0 && http.MethodPut == strings.ToUpper(r.Method) {
			//nameregex = crypt.Base64Dec(nameregex)
			UserPut(nameregex, w, r)
			return
		}
		if endpoint == "/user" && len(nameregex) >= 0 && http.MethodGet == strings.ToUpper(r.Method) {
			//nameregex = crypt.Base64Dec(nameregex)
			UserGet(nameregex, w, r)
			return
		}
		// create url expression, fetch some urls GET
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
