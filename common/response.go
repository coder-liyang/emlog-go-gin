package common

type EmlogResponse struct {
	Error int64       `json:"error"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
}

var responseMap map[int64]string = map[int64]string{
	0: "OK",
}

//统一格式化返回值
func GetResponse(error int64, data interface{}, msg string) EmlogResponse {
	if msg == "" {
		if msg2, ok := responseMap[error]; ok {
			msg = msg2
		} else {
			msg = "未知错误"
		}
	}
	return EmlogResponse{
		Error: 0,
		Data:  data,
		Msg:   msg,
	}
}
