package logic

const (
	NAME        = "小黑"
	PASSWORD    = "密码"
	AGENCY_IP   = "127.0.0.1"
	AGENCY_PORT = "80"
)

// 根据代理找到用户
func FindUserByAgency(ip string, port string) map[string]string {
	UserMap := map[string]string{}

	UserMap["name"] = NAME
	UserMap["password"] = PASSWORD

	return UserMap
}

// 保存用户信息
func saveUser() {
	UserMap := map[string]string{}

	UserMap["name"] = NAME
	UserMap["password"] = PASSWORD
}
