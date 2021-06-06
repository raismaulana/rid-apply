package entity

type User struct {
	Emails []string `json:"emails"`
	Name   struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Sso []struct {
		Provider string `json:"provider"`
		Userid   string `json:"userid"`
	} `json:"sso"`
}
