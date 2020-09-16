package handler

import (
	"net/http"
	"strings"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user/api.user/repo/user"
)

func GetUser(userUuid string, w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == http.MethodGet {
		handlerGetUser(userUuid, w, r)
	} else {
		jsonstr := `{"msg":"O método permitido é Get!"}`
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonstr))
		return
	}
}

func handlerGetUser(userUuid string, w http.ResponseWriter, r *http.Request) {
	content_type := strings.ToLower(r.Header.Get("Content-Type"))
	if content_type != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		jsonstr := `{"msg":"Content-type é obrigatório!"}`
		w.Write([]byte(jsonstr))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user.List(userUuid)))
}
