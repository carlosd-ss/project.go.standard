package handler

import (
	"net/http"
	"strings"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/internal/fmts"
	repo "github.com/jeffotoni/project.go.standard/project-web/standard/crud.user.singleton/repo/user"
)

//UserGet ..
func UserGet(id string, w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == http.MethodGet {
		handlerGetUser(id, w, r)
	} else {
		w.Header().Add("Content-Type", "application/json")
		jsonstr := `{"msg":"O método permitido é Get!"}`
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonstr))
		return
	}
}

func handlerGetUser(id string, w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	/*
		contentType := strings.ToLower(r.Header.Get("Content-Type"))
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			jsonstr := `{"msg":"Content-type é obrigatório!"}`
			w.Write([]byte(jsonstr))
			return
		}
	*/
	w.WriteHeader(http.StatusOK)
	name, lastname, id := repo.List(id)
	w.Write([]byte(fmts.Concat(`{"user":"`, name, `","lastname":"`, lastname, `","id":`, id, `}`)))
	return
}
