package model

type User struct {
	Id       int    `orm:"id,primary" json:"id"`
	Username string `orm:"username" json:"username"`
	Password string `json:"password"`
}
