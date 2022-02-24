package entity

type Result struct {
	Code int
	Msg  string
	Data map[string]interface{}
}

const (
	CodeError = 10000
)

func (r *Result) SetCode(code int) {
	r.Code = code
}

func (r *Result) SetMessage(msg string) {
	r.Msg = msg
}

func (r *Result) SetData(data map[string]interface{}) {
	r.Data = data
}
