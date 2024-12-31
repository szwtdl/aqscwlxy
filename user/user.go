package user

import (
	"github/szwtdl/aqscwlxy/types"
	"github/szwtdl/aqscwlxy/utils"
)

// 账号登录

func Login(client *utils.HttpClient, data map[string]string, header map[string]string) (types.User, error) {
	sign := utils.GetSign(map[string]interface{}{
		"account_number": data["username"],
		"password":       data["password"],
		"platform":       data["platform"],
		"zx_code":        "",
		"t":              data["timestamp"],
		"imei":           data["imei"],
	})
	response, err := client.DoPost("login_account.php", header, map[string]string{
		"account_number": data["username"],
		"password":       data["password"],
		"platform":       data["platform"],
		"zx_code":        "",
		"t":              data["timestamp"],
		"imei":           data["imei"],
		"sign":           sign,
	})
	if err != nil {
		return types.User{}, err
	}
	var responseApi types.ResponseApi
	err = utils.JsonUnmarshal(response, &responseApi)
	if err != nil {
		return types.User{}, err
	}
	var user types.User
	err = utils.JsonUnmarshal(utils.JsonMarshal(responseApi.Data), &user)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

// 检查是否认证,如果是302，代表需要认证

func CheckAuth(client *utils.HttpClient, data map[string]string, header map[string]string) (types.ResponseApi, error) {
	sign := utils.GetSign(map[string]interface{}{
		"platform": data["platform"],
		"zx_code":  "",
		"id":       data["user_id"],
		"imei":     data["imei"],
		"t":        data["timestamp"],
	})
	postData := map[string]string{
		"platform": data["platform"],
		"zx_code":  "",
		"id":       data["user_id"],
		"imei":     data["imei"],
		"t":        data["timestamp"],
		"sign":     sign,
		"token":    data["token"],
	}
	response, err := client.DoPost("user/check_user_auth.php", header, postData)
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

// 检查token，是否过期

func CheckToken(client *utils.HttpClient, data map[string]string, header map[string]string) (types.ResponseApi, error) {
	sign := utils.GetSign(map[string]interface{}{
		"platform": data["platform"],
		"zx_code":  "",
		"id":       data["user_id"],
		"imei":     data["imei"],
		"t":        data["timestamp"],
	})
	response, err := client.DoPost("user/checkToken.php", header, map[string]string{
		"platform": data["platform"],
		"zx_code":  "",
		"id":       data["user_id"],
		"imei":     data["imei"],
		"t":        data["timestamp"],
		"sign":     sign,
		"token":    data["token"],
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
