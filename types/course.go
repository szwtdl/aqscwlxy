package types

type Course struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	TotalHour          string `json:"totalhour"`
	BeginTime          string `json:"begintime"`
	EndTime            string `json:"endtime"`
	Notes              string `json:"learning_notes"`
	CSTotalHour        string `json:"cs_totalhour"`
	CName              string `json:"c_name"`
	CourseId           string `json:"course_id"`
	CDuration          string `json:"c_duration"`
	ClassStudentId     string `json:"classstudent_id"`
	Achievement        string `json:"achievement"`
	IsExam             string `json:"is_exam"`
	StudentName        string `json:"studentname"`
	CheckStudy         string `json:"check_study"`
	TypeName           string `json:"typename"`
	IsAuthFace         string `json:"is_auth_face"`
	AdminId            string `json:"admin_id"`
	IsAuthSub          string `json:"is_auth_sub"`
	OCType             string `json:"oc_type"`
	ExamNew            string `json:"exam_new"`
	IsComplete         string `json:"is_complete"`
	ExamCount          string `json:"exam_count"`
	IsQuestionnaire    string `json:"is_questionnaire"`
	QuestionnaireUrl   string `json:"questionnaire_url"`
	QuestionnaireNum   string `json:"questionnaire_num"`
	PassingGrade       string `json:"passing_grade"`
	ExamProgress       string `json:"exam_progress"`
	CoverPic           string `json:"cover_pic"`
	OperationType      string `json:"operation_type"`
	LearnedTotalHour   string `json:"learned_totalhour"`
	IsPractice         string `json:"is_practice"`
	PracticeNum        string `json:"practice_num"`
	IsAuthentication   string `json:"is_authentication"`
	FPCardPassProgress string `json:"fpcard_pass_progress"`
	SpeedOfProgress    string `json:"speed_of_progress"`
}

type Section struct {
	Id                  string `json:"id"`
	Fid                 string `json:"fid"`
	Name                string `json:"name"`
	Videofile           string `json:"videofile"`
	CourseId            string `json:"course_id"`
	TotalHour           string `json:"totalhour"`
	Introduce           string `json:"introduce"`
	Createtime          string `json:"createtime"`
	Duration            string `json:"duration"`
	SrDuration          string `json:"sr_duration"`
	SrTotalHour         string `json:"sr_totalhour"`
	VideoDuration       string `json:"video_duration"`
	Speaker             string `json:"speaker"`
	VideoCover          string `json:"video_cover"`
	Type                string `json:"type"`
	PreventionCheatTime string `json:"prevention_cheat_time"`
	VideoPathGao        string `json:"video_path_gao"`
	VideoPathBiao       string `json:"video_path_biao"`
	VideoPathLiu        string `json:"video_path_liu"`
	SubjectQuantity     string `json:"subject_quantity"`
	StudyRecordId       string `json:"studyrecord_id"`
	SrStartTime         string `json:"sr_starttime"`
	SrFinishTime        string `json:"sr_finishtime"`
	SrStatus            string `json:"sr_status"`
	SpeedOfProgress     string `json:"speed_of_progress"`
	SrDurationStr       string `json:"sr_duration_str"`
	CvDurationStr       string `json:"cv_duration_str"`
	ValidphotoTime      string `json:"validphoto_time"`
	IsReset             string `json:"is_reset"`
	IsEvaluate          string `json:"is_evaluate"`
	AuthFaceCount       string `json:"auth_face_count"`
	AuthVideoCount      string `json:"auth_video_count"`
	SubAuthStatus       string `json:"sub_auth_status"`
	AuthSubCount        string `json:"auth_sub_count"`
	IsComplete          string `json:"is_complete"`
}

type Chapter struct {
	Id                  string    `json:"id"`
	Fid                 string    `json:"fid"`
	Name                string    `json:"name"`
	Videofile           string    `json:"videofile"`
	Weigh               string    `json:"weigh"`
	CourseId            string    `json:"course_id"`
	TotalHour           string    `json:"totalhour"`
	Introduce           string    `json:"introduce"`
	Createtime          string    `json:"createtime"`
	Duration            string    `json:"duration"`
	SrDuration          int64     `json:"sr_duration"`
	SrTotalHour         string    `json:"sr_totalhour"`
	VideoDuration       int64     `json:"video_duration"`
	Speaker             string    `json:"speaker"`
	VideoCover          string    `json:"video_cover"`
	Type                string    `json:"type"`
	PreventionCheatTime string    `json:"prevention_cheat_time"`
	VideoPathGao        string    `json:"video_path_gao"`
	VideoPathBiao       string    `json:"video_path_biao"`
	VideoPathLiu        string    `json:"video_path_liu"`
	SubjectQuantity     string    `json:"subject_quantity"`
	List                []Section `json:"list"`
	SpeedOfProgress     string    `json:"speed_of_progress"`
}
