package study

import (
	"encoding/json"
	"errors"
	"fmt"
	"github/szwtdl/aqscwlxy/types"
	"github/szwtdl/aqscwlxy/utils"
)

// 章节列表

func ChapterList(client *utils.HttpClient, data map[string]string, header map[string]string) ([]types.Chapter, error) {
	post := map[string]interface{}{
		"platform":        data["platform"],
		"zx_code":         "",
		"id":              data["user_id"],
		"classstudent_id": data["course_id"],
		"t":               data["timestamp"],
		"imei":            data["imei"],
	}
	sign := utils.GetSign(post)
	response, err := client.DoPost("study/study_video.php", header, map[string]string{
		"platform":        data["platform"],
		"zx_code":         "",
		"id":              data["user_id"],
		"t":               data["timestamp"],
		"classstudent_id": data["course_id"],
		"imei":            data["imei"],
		"sign":            sign,
		"token":           data["token"],
	})
	if err != nil {
		return nil, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return nil, err
	}
	var chapterList []types.Chapter
	err = utils.JsonUnmarshal(utils.JsonMarshal(responseApi.Data), &chapterList)
	if err != nil {
		return nil, err
	}
	return chapterList, nil
}

// 章节详情

func ChapterInfo(client *utils.HttpClient, data map[string]string, header map[string]string) (types.Section, error) {
	post := map[string]interface{}{
		"platform":        data["platform"],
		"zx_code":         "",
		"id":              data["user_id"],
		"classstudent_id": data["course_id"],
		"imei":            data["imei"],
		"t":               data["timestamp"],
	}
	sign := utils.GetSign(post)
	response, err := client.DoPost("study/study_video.php", header, map[string]string{
		"platform":        data["platform"],
		"zx_code":         "",
		"id":              data["user_id"],
		"t":               data["timestamp"],
		"imei":            data["imei"],
		"classstudent_id": data["course_id"],
		"sign":            sign,
		"token":           data["token"],
	})
	if err != nil {
		return types.Section{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.Section{}, errors.New("解析响应失败")
	}
	var chapterList []types.Chapter
	err = utils.JsonUnmarshal(utils.JsonMarshal(responseApi.Data), &chapterList)
	if err != nil {
		return types.Section{}, errors.New("解析章节失败")
	}
	for _, chapter := range chapterList {
		for _, section := range chapter.List {
			SpeedOfProgress, _ := utils.RemovePercentAndConvert(section.SpeedOfProgress)
			if SpeedOfProgress < 100 {
				return section, nil
			}
		}
	}
	return types.Section{}, nil
}

// 需要考试的章节

func NeedExamChapter(client *utils.HttpClient, data map[string]string, header map[string]string) (types.Section, error) {
	sign := utils.GetSign(map[string]interface{}{
		"platform":        data["platform"],
		"zx_code":         "",
		"classstudent_id": data["course_id"],
		"id":              data["user_id"],
		"t":               data["timestamp"],
		"imei":            data["imei"],
	})
	response, err := client.DoPost("study/study_video.php", header, map[string]string{
		"platform":        data["platform"],
		"zx_code":         "",
		"classstudent_id": data["course_id"],
		"id":              data["user_id"],
		"t":               data["timestamp"],
		"imei":            data["imei"],
		"sign":            sign,
		"token":           data["token"],
	})
	if err != nil {
		return types.Section{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.Section{}, err
	}
	var chapterList []types.Chapter
	err = utils.JsonUnmarshal(utils.JsonMarshal(responseApi.Data), &chapterList)
	if err != nil {
		return types.Section{}, err
	}
	for _, chapter := range chapterList {
		for _, section := range chapter.List {
			SpeedOfProgress, _ := utils.RemovePercentAndConvert(section.SpeedOfProgress)
			if SpeedOfProgress == 100 && section.IsComplete == "0" {
				return section, nil
			}
		}
	}
	return types.Section{}, nil
}

// 题目列表

func ExamList(client *utils.HttpClient, data map[string]string, header map[string]string) ([]types.Exam, error) {
	post := map[string]interface{}{
		"platform":       data["platform"],
		"zx_code":        "",
		"studyrecord_id": data["record_id"],
		"id":             data["user_id"],
		"t":              data["timestamp"],
		"imei":           data["imei"],
	}
	sign := utils.GetSign(post)
	response, err := client.DoPost("study/get_course_video_subject.php", header, map[string]string{
		"platform":       data["platform"],
		"zx_code":        "",
		"studyrecord_id": data["record_id"],
		"id":             data["user_id"],
		"t":              data["timestamp"],
		"imei":           data["imei"],
		"sign":           sign,
		"token":          data["token"],
	})
	if err != nil {
		return nil, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return nil, err
	}
	var examList []types.Exam
	err = utils.JsonUnmarshal(utils.JsonMarshal(responseApi.Data), &examList)
	if err != nil {
		return nil, err
	}
	return examList, nil
}

// 提交题目

func SubmitExam(client *utils.HttpClient, data map[string]string, header map[string]string) (types.ResponseApi, error) {
	sign := utils.GetSign(map[string]interface{}{
		"studyrecord_id": data["record_id"],
		"platform":       data["platform"],
		"data":           data["data"],
		"zx_code":        "",
		"id":             data["user_id"],
		"imei":           data["imei"],
		"t":              data["timestamp"],
	})
	postData := map[string]string{
		"studyrecord_id": fmt.Sprintf("%v", data["record_id"]),
		"platform":       fmt.Sprintf("%v", data["platform"]),
		"zx_code":        "",
		"id":             fmt.Sprintf("%v", data["user_id"]),
		"imei":           fmt.Sprintf("%v", data["imei"]),
		"t":              fmt.Sprintf("%v", data["timestamp"]),
		"sign":           sign,
		"token":          fmt.Sprintf("%v", data["token"]),
	}

	// 解析 data["data"] 为 []map[string]string
	var nestedData []map[string]string
	if err := json.Unmarshal([]byte(data["data"]), &nestedData); err != nil {
		return types.ResponseApi{}, fmt.Errorf("invalid data format: %v", err)
	}

	// 动态生成 data[n][id] 和 data[n][value] 格式
	for i, item := range nestedData {
		postData[fmt.Sprintf("data[%d][id]", i)] = item["id"]
		postData[fmt.Sprintf("data[%d][value]", i)] = item["value"]
	}

	response, err := client.DoPost("study/subject_complete.php", header, postData)
	if err != nil {
		return types.ResponseApi{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.ResponseApi{}, err
	}
	return responseApi, nil
}

// 视频进度

func SwitchVideo(client *utils.HttpClient, data map[string]string, header map[string]string) (types.ResponseApi, error) {
	post := map[string]interface{}{
		"studyrecord_id": data["record_id"],
		"duration":       data["duration"],
		"platform":       data["platform"],
		"zx_code":        "",
		"id":             data["user_id"],
		"imei":           data["imei"],
		"t":              data["timestamp"],
	}
	sign := utils.GetSign(post)
	response, err := client.DoPost("study/study_video_file.php", header, map[string]string{
		"studyrecord_id": data["record_id"],
		"duration":       data["duration"],
		"platform":       data["platform"],
		"zx_code":        "",
		"id":             data["user_id"],
		"imei":           data["imei"],
		"t":              data["timestamp"],
		"sign":           sign,
		"token":          data["token"],
	})
	if err != nil {
		return types.ResponseApi{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.ResponseApi{}, err
	}
	return responseApi, nil
}

// 提交学习记录

func SubmitRecord(client *utils.HttpClient, data map[string]string, header map[string]string) (types.ResponseApi, error) {
	post := map[string]interface{}{
		"platform":       data["platform"],
		"zx_code":        "",
		"id":             data["user_id"],
		"studyrecord_id": data["record_id"],
		"sr_totalhour":   data["progress"],
		"imei":           data["imei"],
		"t":              data["timestamp"],
	}
	sign := utils.GetSign(post)
	response, err := client.DoPost("study/studyrecord.php", header, map[string]string{
		"platform":       data["platform"],
		"zx_code":        "",
		"id":             data["user_id"],
		"studyrecord_id": data["record_id"],
		"sr_totalhour":   data["progress"],
		"imei":           data["imei"],
		"t":              data["timestamp"],
		"sign":           sign,
		"token":          data["token"],
	})
	if err != nil {
		return types.ResponseApi{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.ResponseApi{}, err
	}
	return responseApi, nil
}
