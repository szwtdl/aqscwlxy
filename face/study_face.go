package face

import (
	"fmt"
	"github/szwtdl/aqscwlxy/types"
	"github/szwtdl/aqscwlxy/utils"
)

// 人脸识别认证

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
	postData := map[string]string{
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
	}
	response, err := client.DoPost("study/face_auth.php", postData)
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
