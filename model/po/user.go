package po

type LoginType int

type User struct {
	Id       uint64 `json:"id" gorm:"column:id"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Salt     string `json:"salt" gorm:"column:salt"`
	Passw    string `json:"passw" gorm:"column:passw"`
	Birthday string `json:"birthday" gorm:"column:birthday"`
	Email    string `json:"email" gorm:"column:email"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Sex      int    `json:"sex" gorm:"column:sex"`
	Address  string `json:"address" gorm:"column:address"`
	Ctime    int    `json:"ctime" gorm:"column:ctime"`
	Mtime    int    `json:"mtime" gorm:"column:mtime"`
	Taste    int    `json:"taste" gorm:"column:taste"`
}

func (p *User) TableName() string {
	return "user_tab"
}

var UserTab = new(User)
