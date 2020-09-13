// Go Api server
// @jeffotoni

package controller

type EndpointS struct {
	Ping string
	User string
}

func Endpoint() *EndpointS {
	return &EndpointS{
		Ping: "/user/ping", // POST/GET
		User: "/user",      // [DELETE] [GET]
	}
}
