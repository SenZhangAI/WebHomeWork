package main

type User struct {
	ID       int32  `form:"user_id" json:"user_id"`
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	UserNick string `form:"user_nick" json:"user_nick" binding:"required"`
	Password string `form:"user_password" json:"user_password" binding:"required"`
}
