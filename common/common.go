package common

type EmlogResponse struct {
	Error int64       `json:"error"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
}

func GetResponse (error int64, data interface{}, msg string) EmlogResponse {
	return EmlogResponse{
		Error: 0,
		Data: data,
		Msg: "OK",
	}
}