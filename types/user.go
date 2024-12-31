package types

type User struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Account          string `json:"account"`
	Mobile           string `json:"mobile"`
	IdCard           string `json:"idcard"`
	Gender           string `json:"gender"`
	IsAuthentication string `json:"is_authentication"`
	Head             string `json:"head"`
	Token            string `json:"token"`
	IsSetPassword    bool   `json:"is_set_password"`
	UserType         string `json:"user_type"`
	CreateTime       string `json:"createtime"`
	IsValid          string `json:"isvalid"`
}
