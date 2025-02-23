package org

import (
	"errors"
	"github.com/szwtdl/aqscwlxy/types"
	"github.com/szwtdl/aqscwlxy/utils"
	"github.com/szwtdl/req"
)

func CourseList(client *client.HttpClient, data map[string]string) ([]types.Course, error) {
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
	if responseApi.Code != 200 {
		return nil, errors.New(responseApi.Msg)
	}
	var courseList []types.Course
	err = utils.JsonUnmarshal(utils.JsonMarshal(responseApi.Data), &courseList)
	if err != nil {
		return nil, err
	}
	return courseList, nil
}

// 课程详情

func CourseInfo(client *client.HttpClient, data map[string]string) (types.CourseDetail, error) {
	post := map[string]interface{}{
		"platform":        data["platform"],
		"zx_code":         "",
		"id":              data["user_id"],
		"t":               data["timestamp"],
		"imei":            data["imei"],
		"classstudent_id": data["course_id"],
	}
	sign := utils.GetSign(post)
	response, err := client.DoPost("org_class/get_orgclass_info.php", map[string]string{
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
		return types.CourseDetail{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if responseApi.Code != 200 {
		return types.CourseDetail{}, errors.New(responseApi.Msg)
	}
	if err != nil {
		return types.CourseDetail{}, err
	}
	var courseDetail types.CourseDetail
	err = utils.JsonUnmarshal(utils.JsonMarshal(responseApi.Data), &courseDetail)
	if err != nil {
		return types.CourseDetail{}, err
	}
	return courseDetail, nil
}
