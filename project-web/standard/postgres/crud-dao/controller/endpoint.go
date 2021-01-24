package controller

// Go Api server
// @jeffotoni
// @ancogamer

//EndpointS ..
type EndpointS struct {
	Ping string
	User string
}

//Endpoint ..
func Endpoint() *EndpointS {
	return &EndpointS{
		Ping: "/ping", // POST/GET
		User: "/user", // [POST] [GET] [UPDATE] [DELETE]
	}
}
