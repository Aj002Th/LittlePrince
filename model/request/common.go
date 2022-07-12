package request

type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

type UserInfo struct {
	Name string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
}

type IdReq struct {
	ID uint `json:"id" form:"id"`
}
