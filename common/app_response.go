package common

// Response represents a standard API response
type Response struct {
	Data   interface{} `json:"data"`
}

type successRes struct {
	Data   interface{} `json:"data"`
}

func SimpleSuccessRes(data interface{}) *successRes {
    return &successRes{data}
}
