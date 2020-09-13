// Go Api server
// @jeffotoni
// 2019-05-14

package util

import (
	"encoding/json"
	"net/http"
)

type JsonMsgEspec struct {
	Status string
	Msg    string
}

type JsonMsg JsonMsgEspec

// Returns json without typing in http
func WriteJson(Status string, Msg string) {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"msg":"We could not generate the json error!"}`
		Print(msg)
		return
	}

	// byte
	Print(string(msgJson))
}

// Returns json by typing on http
func HttpWriteJson(w http.ResponseWriter, Status string, Msg string, httpStatus int) {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"msg":"We could not generate the json error!"}`
		w.WriteHeader(http.StatusForbidden)
		//io.WriteString(w, msg)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msgJson)
}

// Returns json by typing on http
func GetJson(w http.ResponseWriter, Status string, Msg string, httpStatus int) string {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"msg":"We could not generate the json error!"}`
		return msg
	}
	return string(msgJson)
}

func MsgErrorJson(msg string) string {
	return BuildJSON(`{"msg":"`, msg, `"}`)
}

func DoJson(Status string, Msg string) string {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"msg":"We could not generate the json error!"}`
		return msg
	}
	return string(msgJson)
}
