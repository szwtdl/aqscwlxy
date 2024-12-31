package main

import (
	"fmt"
	"github/szwtdl/aqscwlxy/study"
	"github/szwtdl/aqscwlxy/user"
	"github/szwtdl/aqscwlxy/utils"
)

var httpClient = utils.NewClient()
var deviceImei = utils.GetUuid("2")

func main() {
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
	}
	platformType := "1"
	loginPost := map[string]string{
		"username":  "41142219890415121X",
		"password":  "15121X",
		"platform":  platformType,
		"timestamp": utils.GetReqTime(),
		"imei":      deviceImei,
	}
	User, err := user.Login(httpClient, loginPost, header)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println("登录账号:", User.Id, User.Name, User.Account, User.Token, deviceImei, utils.GetRandomFloat())
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
	//response, err := user.CheckAuth(httpClient, checkPost, header)
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
	//	}, header)
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
	//	"user_id":   User.Id,
	//	"token":     User.Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//	"course_id": "240746",
	//}, header)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//if chapterInfo.StudyRecordId == "" {
	//	fmt.Println("Chapter Info StudyRecordId is empty")
	//	return
	//}
	//switchVideo, err := study.SwitchVideo(httpClient, map[string]string{
	//	"user_id":   User.Id,
	//	"token":     User.Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//	"record_id": chapterInfo.StudyRecordId,
	//}, header)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//fmt.Println("切换视频状态:", switchVideo.Code)
	//progressTime, _ := strconv.Atoi(chapterInfo.SrDuration)
	//duration, _ := strconv.Atoi(chapterInfo.Duration)
	//fmt.Println("视频进度", progressTime, duration)
	//// 睡眠5秒
	//utils.Sleep(60)
	//response, err := study.SubmitRecord(httpClient, map[string]string{
	//	"user_id":   User.Id,
	//	"token":     User.Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//	"record_id": chapterInfo.StudyRecordId,
	//	"progress":  strconv.Itoa(progressTime + 60),
	//}, header)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//fmt.Println(response.Msg, response.Code)

	// TODO 查询需要考试的章节
	section, err := study.NeedExamChapter(httpClient, map[string]string{
		"user_id":   User.Id,
		"token":     User.Token,
		"platform":  platformType,
		"imei":      deviceImei,
		"timestamp": utils.GetReqTime(),
		"course_id": "240746",
	}, header)
	if err != nil {
		println(err.Error())
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
	}, header)
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
	examBody, err := study.SubmitExam(httpClient, examPost, header)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println("提交题目结果:", examBody.Code, examBody.Msg)
	//coursePost := map[string]string{
	//	"user_id":   User.Id,
	//	"token":     User.Token,
	//	"platform":  platformType,
	//	"imei":      deviceImei,
	//	"timestamp": utils.GetReqTime(),
	//}
	//////获取课程列表
	//courseList, err := org.CourseList(httpClient, coursePost, header)
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
	//	}, header)
	//	if err != nil {
	//		println(err.Error())
	//		return
	//	}
	//	fmt.Println("Course Info:", chapterInfo.Name, chapterInfo.StudyRecordId, chapterInfo.SrDuration, chapterInfo.Duration)
	//}

}
