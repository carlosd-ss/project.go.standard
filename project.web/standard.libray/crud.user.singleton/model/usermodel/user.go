package usermodel

type User struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Age      int    `json:"age,omitempty"`
}
