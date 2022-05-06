package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, psw string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: psw}).First(&auth)

	if auth.ID > 0 {
		return true
	}
	return false
}
