package serializer

type Response struct {
	Code  int         `json:"status"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
}

type TokenResponse struct {
	token string
}
