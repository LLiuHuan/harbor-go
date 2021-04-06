package schema

type CurrencyError struct {
	ERROR []Response
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

func Result(code int, data interface{}, msg interface{}) Response {
	return Response{
		Code:    code,
		Data:    data,
		Message: msg,
	}
}

func OkWithMessage(message interface{}) Response {
	return Result(200, map[string]interface{}{}, message)
}

func OkWithData(data interface{}) Response {
	return Result(200, data, "操作成功")
}

func OkWithDetailed(code int, data interface{}, message interface{}) Response {
	return Result(code, data, message)
}
