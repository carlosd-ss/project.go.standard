// Go Api server
// @jeffotoni
// 2019-05-14

package middleware

//import (
//"net/http"

//"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/gjwt"
//)

//func AutJwt() Adapter {
//return func(h http.Handler) http.Handler {
//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//var ok bool = false
//var jsonerr string = `{"msg":"token-invalid"}`
//ok = gjwt.GtokenJwt2(w, r)
//if ok {
//h.ServeHTTP(w, r)
//} else {
//w.Header().Set("Content-Type", "application/json")
//w.WriteHeader(http.StatusUnauthorized)
//w.Write([]byte(jsonerr))
//return
//}
//})
//}
//}
