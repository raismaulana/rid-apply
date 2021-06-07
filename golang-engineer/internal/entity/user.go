package entity

type User struct {
	ID     int64   `gorm:"primary_key:auto_increment" json:"id"`
	Name   string  `gorm:"type:varchar(255)" json:"name"`
	Emails []Email `gorm:"foreignKey:UserId;references:ID" json:"emails"`
	Sso    []SSO   `gorm:"foreignKey:UserId;references:ID" json:"sso"`
}

type Email struct {
	ID     int64  `gorm:"primary_key:auto_increment" json:"id"`
	Email  string `gorm:"type:varchar(255) unique" json:"email"`
	UserId int64  `json:"user_id"`
}

type SSO struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"id"`
	Provider string `gorm:"type:varchar(255)" json:"provider"`
	Userid   string `gorm:"type:varchar(255) unique" json:"userid"`
	UserId   int64  `json:"user_id"`
}

type Google struct {
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	ID            string `json:"id"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
