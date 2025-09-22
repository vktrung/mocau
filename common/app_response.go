package common

// Response represents a standard API response
type Response struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessRes(data interface{}, paging interface{}, filter interface{}) *successRes {
	return &successRes{data, paging, filter}
}

func SimpleSuccessRes(data interface{}) *successRes {
	return &successRes{data, nil, nil}
}
