package org

import (
	"github.com/szwtdl/aqscwlxy/types"
	"github.com/szwtdl/aqscwlxy/utils"
)

func CourseList(client *utils.HttpClient, data map[string]string) ([]types.Course, error) {
	sign := utils.GetSign(map[string]interface{}{
		"platform": data["platform"],
		"zx_code":  "",
		"id":       data["user_id"],
		"t":        data["timestamp"],
		"imei":     data["imei"],
	})
	response, err := client.DoPost("org_class/get_study_list.php", map[string]string{
		"platform": data["platform"],
		"zx_code":  "",
		"id":       data["user_id"],
		"imei":     data["imei"],
		"t":        data["timestamp"],
		"sign":     sign,
		"token":    data["token"],
	})
	if err != nil {
		return nil, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return nil, err
	}
	var courseList []types.Course
	err = utils.JsonUnmarshal(utils.JsonMarshal(responseApi.Data), &courseList)
	if err != nil {
		return nil, err
	}
	return courseList, nil
}

// 课程详情

func CourseInfo(client *utils.HttpClient, data map[string]string) (types.ResponseApi, error) {
	post := map[string]interface{}{
		"platform":        data["username"],
		"zx_code":         "",
		"id":              data["user_id"],
		"t":               data["timestamp"],
		"imei":            data["imei"],
		"classstudent_id": data["course_id"],
	}
	sign := utils.GetSign(post)
	response, err := client.DoPost("org_class/get_orgclass_info.php", map[string]string{
		"platform":        data["username"],
		"zx_code":         "",
		"id":              data["user_id"],
		"t":               data["timestamp"],
		"imei":            data["imei"],
		"classstudent_id": data["course_id"],
		"sign":            sign,
		"token":           data["token"],
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
