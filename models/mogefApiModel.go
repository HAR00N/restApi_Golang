package models

type ResData struct {
	Header []Header `json:"header"`
	Body   []Body   `json:"body"`
}

type Header struct {
	ResultCode int    `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
}

type Body struct {
	PageNo     int     `json:"pageNo"`
	TotalCount int     `json:"totalCount"`
	NumOfRows  int     `json:"numOfRows"`
	Items      []Items `json:"items"`
}

type Items struct {
	RegDt    string `json:"regDt"`
	ThumbUrl string `json:"thumbUrl"`
	Title    string `json:"title"`
	ViewUrl  string `json:"viewUrl"`
}
