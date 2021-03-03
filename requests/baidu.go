package requests

// 公共的header
type Header struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Target   string `json:"target"`
}


type ReportBody struct {
	ReportId string `json:"reportId"`
}

type ReoprtReq struct {
	Header Header     `json:"header"`
	Body   ReportBody `json:"body"`
}

type Params interface {
	toJson() []byte
}
