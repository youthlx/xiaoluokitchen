package vo

type UserInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	BirthDay string `json:"birth_day"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"`
	Address  string `json:"address"`
	Taste    int    `json:"taste"`
}
