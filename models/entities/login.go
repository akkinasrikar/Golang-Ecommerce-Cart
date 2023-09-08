package entities

type SignUp struct {
	UserId   int64  `gorm:"column:users_id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:user_name;not null"`
	Email    string `gorm:"column:user_email;not null;unique"`
	Password string `gorm:"column:password;not null"`
}

type Login struct {
	Name     string `gorm:"column:user_name;not null"`
	Password string `gorm:"column:password;not null"`
}
