package baseresponse

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// func NewSuccessBaseResponse() BaseResponse {
// 	return BaseResponse{
// 		Code:    0,
// 		Message: "Success",
// 	}
// }

// func NewSuccessWithData(data interface{}) BaseResponse {
// 	return BaseResponse{
// 		Code:    0,
// 		Message: "Success",
// 		Data:    data,
// 	}
// }
