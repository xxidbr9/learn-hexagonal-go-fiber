package shared

// Res :
type Res struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type Error struct {
	Error string `json:"error"`
}

// NewRes :
func NewRes(msg string, data interface{}) interface{} {
	return &Res{Msg: msg, Data: data}
}

// ErrorRes : Error JSON Response
func ErrorRes(err error) interface{} {
	return &Error{Error: err.Error()}
}
