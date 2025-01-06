package face

import (
	"fmt"
	"github.com/szwtdl/aqscwlxy/types"
	"github.com/szwtdl/aqscwlxy/utils"
)

// 返回待认证数据

func DurationFace(client *utils.HttpClient, data map[string]string) (types.ResponseApi, error) {
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
	return responseApi, nil
}

//  人脸认证

func StudyFace(client *utils.HttpClient, data map[string]string) (types.ResponseApi, error) {
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
	response, err := client.DoPost("study/face_study.php", map[string]string{
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
		return types.ResponseApi{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.ResponseApi{}, err
	}
	return responseApi, nil
}

// 首次认证头像

func AuthFace(client *utils.HttpClient, data map[string]string) (types.ResponseApi, error) {
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
	return responseApi, nil
}

// 实名认证

func UploadFace(client *utils.HttpClient, data map[string]string) (types.ResponseApi, error) {
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
	return responseApi, nil
}
