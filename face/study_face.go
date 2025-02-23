package face

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/szwtdl/aqscwlxy/types"
	"github.com/szwtdl/aqscwlxy/utils"
	"github.com/szwtdl/req"
)

// 返回待认证数据

func DurationFace(client *client.HttpClient, data map[string]string) (types.ResponseApi, error) {
	sign := utils.GetSign(map[string]interface{}{
		"studyrecord_id": data["record_id"],
		"duration":       data["duration"],
		"platform":       data["platform"],
		"zx_code":        "",
		"id":             data["user_id"],
		"imei":           data["imei"],
		"t":              data["timestamp"],
	})
	response, err := client.DoPost("face/launch_face_auth.php", map[string]string{
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
	if responseApi.Code != 200 {
		return types.ResponseApi{}, fmt.Errorf(responseApi.Msg)
	}
	return responseApi, nil
}

//  人脸认证

func StudyFace(client *client.HttpClient, data map[string]string) (types.ResponseApi, error) {
	sign := utils.GetSign(map[string]interface{}{
		"studyrecord_id":      data["record_id"],
		"base64_image":        data["image"],
		"duration":            data["duration"],
		"studyrecord_face_id": data["face_id"],
		"uuid":                "",
		"zx_code":             "",
		"platform":            data["platform"],
		"id":                  data["user_id"],
		"imei":                data["imei"],
		"t":                   data["timestamp"],
	})
	response, err := client.DoPost("face/study_face.php", map[string]string{
		"studyrecord_id":      data["record_id"],
		"base64_image":        data["image"],
		"duration":            data["duration"],
		"studyrecord_face_id": data["face_id"],
		"uuid":                "",
		"zx_code":             "",
		"platform":            data["platform"],
		"id":                  data["user_id"],
		"imei":                data["imei"],
		"t":                   data["timestamp"],
		"sign":                sign,
		"token":               data["token"],
	})
	if err != nil {
		return types.ResponseApi{}, errors.New("服务端错误")
	}
	var responseApi types.ResponseApi
	err = json.Unmarshal(response, &responseApi)
	if err != nil {
		return types.ResponseApi{}, errors.New("解析数据失败:" + err.Error())
	}
	if responseApi.Code != 200 {
		return types.ResponseApi{}, errors.New("人脸认证失败:" + responseApi.Msg)
	}
	return responseApi, nil
}

// 首次认证头像

func AuthFace(client *client.HttpClient, data map[string]string) (types.ResponseApi, error) {
	sign := utils.GetSign(map[string]interface{}{
		"studyrecord_id":      data["record_id"],
		"base64_image":        data["image"],
		"duration":            data["duration"],
		"studyrecord_face_id": data["face_id"],
		"uuid":                "",
		"zx_code":             "",
		"platform":            "1",
		"id":                  data["user_id"],
		"imei":                data["imei"],
		"t":                   data["timestamp"],
	})
	response, err := client.DoPost("study/face_auth.php", map[string]string{
		"studyrecord_id":      data["record_id"],
		"base64_image":        data["image"],
		"duration":            data["duration"],
		"studyrecord_face_id": data["face_id"],
		"uuid":                "",
		"zx_code":             "",
		"platform":            "1",
		"id":                  data["user_id"],
		"imei":                data["imei"],
		"t":                   data["timestamp"],
		"sign":                sign,
		"token":               data["token"],
	})
	if err != nil {
		return types.ResponseApi{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.ResponseApi{}, err
	}
	if responseApi.Code != 200 {
		return types.ResponseApi{}, fmt.Errorf(responseApi.Msg)
	}
	return responseApi, nil
}

// 实名认证

func UploadFace(client *client.HttpClient, data map[string]string) (types.ResponseApi, error) {
	postData := map[string]string{
		"id":        data["user_id"],
		"token":     data["token"],
		"i":         data["imei"],
		"uname":     data["nickname"],
		"uidnum":    data["username"],
		"imgbase64": data["image"],
		"fr":        "",
		"type":      "1",
		"zx_code":   "",
		"user_type": "1",
	}
	response, err := client.DoPost(fmt.Sprintf("send_check.php?t=%f", utils.GetRandomFloat()), postData)
	if err != nil {
		return types.ResponseApi{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.ResponseApi{}, err
	}
	if responseApi.Code != 200 {
		return types.ResponseApi{}, fmt.Errorf(responseApi.Msg)
	}
	return responseApi, nil
}
