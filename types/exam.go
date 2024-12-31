package types

type Exam struct {
	Id       string `json:"id"`
	TqName   string `json:"tq_name"`
	TqType   string `json:"tq_type"`
	BzAnswer string `json:"bz_answer"`
	Options  string `json:"xx"`
	Analysis string `json:"analysis"`
}
