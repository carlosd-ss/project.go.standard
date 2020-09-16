package gjwt

import (
	"net/http"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/pkg/util"
)

// Returns json without typing in http
func WriteJson(Status string, Msg string) {
	// msgJsonStruct := &util.JsonMsg{Msg}
	// msgJson, errj := json.Marshal(msgJsonStruct)
	// if errj != nil {
	// 	msg := `{"msg":"We could not generate the json error!"}`
	// 	util.Print(msg)
	// 	return
	// }
	util.Print(util.BuildJSON(`{"msg":"`, Msg, `"}`))
	// byte
	//util.Print(string(msgJson))
}

// Returns json by typing on http
func HttpWriteJson(w http.ResponseWriter, Status string, Msg string, httpStatus int) {
	// msgJsonStruct := &util.JsonMsg{Msg}
	// msgJson, errj := json.Marshal(msgJsonStruct)
	// if errj != nil {
	// 	msg := `{"msg":"We could not generate the json error!"}`
	// 	w.WriteHeader(http.StatusForbidden)
	// 	w.Write([]byte(msg))
	// 	return
	// }
	msgJson := util.BuildJSON(`{"msg":"`, Msg, `"}`)
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(msgJson))
}

// Returns json by typing on http
func GetJson(w http.ResponseWriter, Status string, Msg string, httpStatus int) string {
	return util.BuildJSON(`{"msg":"`, Msg, `"}`)
	//msgJsonStruct := &util.JsonMsg{Msg}
	//msgJson, errj := json.Marshal(msgJsonStruct)
	// if errj != nil {
	// 	msg := `{"msg":"We could not generate the json error!"}`
	// 	return msg
	// }
	//return string(msgJson)
}
