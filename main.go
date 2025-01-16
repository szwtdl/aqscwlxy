package main

import (
	"fmt"
	"github.com/szwtdl/aqscwlxy/study"
	"github.com/szwtdl/aqscwlxy/user"
	"github.com/szwtdl/aqscwlxy/utils"
)

var platformType = "1"
var httpClient = utils.NewClient("https://gd.aqscwlxy.com/gd_api", map[string]string{
	"Content-Type": "application/x-www-form-urlencoded",
	"User-Agent":   "Mozilla/5.0 (Linux; U; Android 12.1.1; zh-cn; OPPO R9sk Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36 OppoBrowser/10.5.1.2",
})
var deviceImei = utils.GetUuid(platformType)

func main() {
	loginPost := map[string]string{
		"username":  "41142219890415121X",
		"password":  "Dp111111",
		"platform":  platformType,
		"timestamp": utils.GetReqTime(),
		"imei":      deviceImei,
	}
	User, err := user.Login(httpClient, loginPost)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("登录账号:", User.Id, User.Name, User.Account, User.Token, deviceImei, utils.GetRandomFloat())
	//response, err := org.CourseInfo(httpClient, map[string]string{
	//	"platform":  platformType,
	//	"user_id":   User.Id,
	//	"timestamp": utils.GetReqTime(),
	//	"imei":      deviceImei,
	//	"course_id": "245345",
	//	"token":     User.Token,
	//})
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(response.CourseInfo.Name)

	//response, err := user.ResetPass(httpClient, map[string]string{
	//	"nickname":         User.Name,
	//	"username":         User.Account,
	//	"old_password":     "Dp123456",
	//	"new_password":     "Dp111111",
	//	"confirm_password": "Dp111111",
	//	"platform":         platformType,
	//	"user_id":          User.Id,
	//	"imei":             deviceImei,
	//	"timestamp":        utils.GetReqTime(),
	//	"token":            User.Token,
	//})
	//if err != nil {
	//	fmt.Println("修改密码失败", err.Error())
	//	return
	//}
	//if response.Code != 200 {
	//	fmt.Println("修改失败:", response.Msg)
	//	return
	//}
	//fmt.Println(response.Msg, response.Code)
	// 干星魁 be926e9b8be3dd705b05a686e406d8fc 5148 41142219890415121X 17360903201308881012
	//UserId := "5148"
	//deviceImei = "17360903201308881012"
	//Token := "be926e9b8be3dd705b05a686e406d8fc"
	//postData := map[string]interface{}{
	//	"studyrecord_id": "123",
	//	"data": map[string]interface{}{
	//		"key1":  " value1 ",
	//		"key 2": "value 2",
	//	},
	//	"platform": 1,
	//	"zx_code":  "",
	//	"id":       "456",
	//	"imei":     "789",
	//	"t":        "2024-12-31",
	//}
	//sign := utils.GetSign(postData)
	//fmt.Println("签名结果:", sign)
	//checkPost := map[string]string{
	//	"platform":  platformType,
	//	"user_id":   User.Id,
	//	"timestamp": utils.GetReqTime(),
	//	"imei":      deviceImei,
	//	"token":     User.Token,
	//}
	//// 验证是否第一次人脸
	//response, err := user.CheckAuth(httpClient, checkPost)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//if response.Code == 302 || response.Code == 208 {
	//	fmt.Println("账号未认证，认证后可以学习")
	//	image, err := utils.ImageToBase64("/Users/mac/go/aqscwlxy/demo/output_0005.jpg")
	//	if err != nil {
	//		println(err.Error())
	//		return
	//	}
	//	responseFace, err := face.UploadFace(httpClient, map[string]string{
	//		"user_id":  User.Id,
	//		"token":    User.Token,
	//		"nickname": User.Name,
	//		"username": User.Account,
	//		"image":    fmt.Sprintf("data:image/jpeg;base64,%s", image),
	//		"imei":     deviceImei,
	//	})
	//	if err != nil {
	//		println(err.Error())
	//		return
	//	}
	//	if responseFace.Code != 200 {
	//		fmt.Println(responseFace.Msg, responseFace.Code)
	//		return
	//	}
	//	fmt.Println("认证成功")
	//	return
	//}
	//fmt.Println(response.Msg, response.Code)
	//chapterInfo, err := study.ChapterInfo(httpClient, map[string]string{
	//	"user_id":   "208819",
	//	"token":     "e6ea282a986eed7c3c6ce7a5a4f796b4",
	//	"platform":  platformType,
	//	"imei":      "17364963500309581258",
	//	"timestamp": utils.GetReqTime(),
	//	"course_id": "245275",
	//})
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//if chapterInfo.StudyRecordId == "" {
	//	fmt.Println("章节信息为空")
	//	return
	//}
	//fmt.Println("Chapter Info:", chapterInfo.Name, chapterInfo.StudyRecordId, chapterInfo.Duration, chapterInfo.SrDuration, chapterInfo.SpeedOfProgress)
	//switchVideo, err := study.SwitchVideo(httpClient, map[string]string{
	//	"user_id":   User.Id,
	//	"token":     User.Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//	"record_id": chapterInfo.StudyRecordId,
	//})
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//var data struct {
	//	AuthFaceCount     int   `json:"auth_face_count"`
	//	Duration          int   `json:"duration"`
	//	ErrorCode         int   `json:"error_code"`
	//	StudyRecordFaceId int64 `json:"studyrecord_face_id"`
	//	StudyRecordId     int64 `json:"studyrecord_id"`
	//}
	//err = utils.JsonUnmarshal(utils.JsonMarshal(switchVideo.Data), &data)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//duration, _ := strconv.Atoi(chapterInfo.Duration)
	//srDuration, _ := strconv.Atoi(chapterInfo.SrDuration)
	//fmt.Println("视频进度", srDuration, duration, chapterInfo.ValidphotoTime)
	//items := utils.GenerateRules(duration, 900, data.Duration, false, "photo", "photo")
	//fmt.Println(items)
	////if items.Rules[items.Index].Time < srDuration {
	////	items.Index++
	////}
	//fmt.Println("生成规则:", items.Rules[items.Index].Time, items.Rules[items.Index].Type)
	//fmt.Println("切换视频状态:", switchVideo.Code, switchVideo.Msg, data.AuthFaceCount, data.Duration, data.ErrorCode, data.StudyRecordFaceId, data.StudyRecordId)
	//progressTime, _ := strconv.Atoi(chapterInfo.SrDuration)
	//duration, _ := strconv.Atoi(chapterInfo.Duration)
	//fmt.Println("视频进度", progressTime, duration)
	//// 睡眠5秒
	//utils.Sleep(60)
	//switchVideo, err := study.SwitchVideo(httpClient, map[string]string{
	//	"user_id":   UserId,
	//	"token":     Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//	"record_id": chapterInfo.StudyRecordId,
	//	"duration":  chapterInfo.SrDuration,
	//})
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//fmt.Println("切换视频状态:", switchVideo.Code, switchVideo.Msg)
	//responseRecord, err := study.SubmitRecord(httpClient, map[string]string{
	//	"user_id":   UserId,
	//	"token":     Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//	"record_id": chapterInfo.StudyRecordId,
	//	"duration":  "360",
	//})
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//fmt.Println("提交学习记录:", responseRecord.Code, responseRecord.Msg)

	// TODO 提交学习进度
	//response, err := face.DurationFace(httpClient, map[string]string{
	//	"user_id":   User.Id,
	//	"token":     User.Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//	"record_id": chapterInfo.StudyRecordId,
	//	"duration":  chapterInfo.SrDuration,
	//})
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//var faceDuration types.FaceDuration
	//err = utils.JsonUnmarshal(utils.JsonMarshal(response.Data), &faceDuration)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//Base64Str, err := utils.ImageToBase64("/Users/mac/go/aqscwlxy/demo/output_0005.jpg")
	//if err != nil {
	//	fmt.Println("转换图片失败", err.Error())
	//	return
	//}
	//fmt.Println("人脸认证位置:", faceDuration.StudyRecordId, faceDuration.Duration, faceDuration.AuthFaceCount, faceDuration.IsVerification, faceDuration.StudyRecordFaceId)
	//response, err = face.StudyFace(httpClient, map[string]string{
	//	"user_id":   User.Id,
	//	"token":     User.Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//	"record_id": chapterInfo.StudyRecordId,
	//	"image":     fmt.Sprintf("data:image/jpeg;base64,%s", Base64Str),
	//	"duration":  chapterInfo.SrDuration,
	//	"face_id":   fmt.Sprintf("%d", faceDuration.StudyRecordFaceId),
	//})
	//if err != nil {
	//	fmt.Println("人脸认证失败", err.Error())
	//	return
	//}
	//fmt.Println(response.Msg, response.Code)
	//if response.Code == 208 {
	//	if response.Msg == "该视频已看完，请按要求继续学习！" {
	//		fmt.Println("视频已看完")
	//		return
	//	}
	//	if response.Msg == "更新视频进度失败，认证图片为空！" {
	//		fmt.Println("认证图片为空")
	//		return
	//	}
	//	if response.Msg == "更新失败，非法提交学习进度，请刷新后重新观看！" {
	//		fmt.Println("非法提交学习进度")
	//		return
	//	}
	//}
	//if response.Code == 301 {
	//	fmt.Println("账号已退出登录")
	//	return
	//}
	//err = utils.JsonUnmarshal(utils.JsonMarshal(response.Data), &faceDuration)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//fmt.Println("人脸认证位置:", faceDuration.StudyRecordId, faceDuration.Duration, faceDuration.AuthFaceCount, faceDuration.IsVerification, faceDuration.StudyRecordFaceId)
	// TODO 查询需要考试的章节
	section, err := study.NeedExamChapter(httpClient, map[string]string{
		"user_id":   User.Id,
		"token":     User.Token,
		"platform":  platformType,
		"imei":      deviceImei,
		"timestamp": utils.GetReqTime(),
		"course_id": "240746",
	})
	if err != nil {
		println(err.Error())
		return
	}
	if section.StudyRecordId == "" {
		fmt.Println("章节信息为空")
		return
	}
	fmt.Println("章节信息:", section.StudyRecordId, section.Name)
	// 获取考试题目列表
	examList, err := study.ExamList(httpClient, map[string]string{
		"user_id":   User.Id,
		"token":     User.Token,
		"platform":  platformType,
		"imei":      deviceImei,
		"timestamp": utils.GetReqTime(),
		"record_id": section.StudyRecordId,
	})
	if err != nil {
		println(err.Error())
		return
	}
	// 声明一个切片，用于存储提交的题目
	items := utils.GetSubject(examList)
	examPost := map[string]string{
		"user_id":   User.Id,
		"token":     User.Token,
		"platform":  platformType,
		"imei":      deviceImei,
		"timestamp": utils.GetReqTime(),
		"record_id": section.StudyRecordId,
		"data":      string(utils.JsonMarshal(items)),
	}
	examBody, err := study.SubmitExam(httpClient, examPost)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println("提交考试结果:", examBody.Code, examBody.Msg)
	//coursePost := map[string]string{
	//	"user_id":   User.Id,
	//	"token":     User.Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//}
	//////获取课程列表
	//courseList, err := org.CourseList(httpClient, coursePost)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//for _, course := range courseList {
	//	fmt.Println("Course Id:", course.ClassStudentId)
	//	fmt.Println("Course Name:", course.CName)
	//	fmt.Println("Course cover pic:", course.CoverPic)
	//	fmt.Println("Course endTime:", course.EndTime)
	//	fmt.Println("Course progress:", course.SpeedOfProgress)
	//	chapterInfo, err := study.ChapterInfo(httpClient, map[string]string{
	//		"user_id":   User.Id,
	//		"token":     User.Token,
	//		"platform":  platformType,
	//		"imei":      deviceImei,
	//		"timestamp": utils.GetReqTime(),
	//	})
	//	if err != nil {
	//		println(err.Error())
	//		return
	//	}
	//	fmt.Println("Course Info:", chapterInfo.Name, chapterInfo.StudyRecordId, chapterInfo.SrDuration, chapterInfo.Duration)
	//}

}
