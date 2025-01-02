package types

type ResponseApi struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Rule struct {
	Time   int    `json:"time"`
	Type   string `json:"type,omitempty"`
	Action string `json:"action,omitempty"`
}

type RulesData struct {
	Rules     []Rule `json:"rules"`
	NextRule  *Rule  `json:"next_rule"`
	BeginTime int    `json:"begin_time"`
	Index     int    `json:"index"`
	Spacing   int    `json:"spacing"`
}
