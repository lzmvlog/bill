package model

type User struct {
	Id       int    `orm:"id,primary" json:"id"`
	UserName string `orm:"username" json:"username"`
	PassWord string `json:"password"`
}
