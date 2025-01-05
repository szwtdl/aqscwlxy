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

type FaceDuration struct {
	ErrorCode         int   `json:"error_code"`
	StudyRecordId     int64 `json:"studyrecord_id"`
	Duration          int   `json:"duration"`
	AuthFaceCount     int   `json:"auth_face_count"`
	AuthStartTime     int   `json:"auth_start_time"`
	IsVerification    bool  `json:"is_verification"`
	StudyRecordFaceId int64 `json:"studyrecord_face_id"`
}
