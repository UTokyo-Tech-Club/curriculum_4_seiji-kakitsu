package model

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RequestDataType struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var RequestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
