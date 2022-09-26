package serializer

type Response struct {
	Code  int         `json:"status"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
}

type TokenData struct {
	User          interface{} `json:"user"`
	Access_token  string      `json:"access_token"`
	Refresh_token string      `json:"refresh_token"`
}
