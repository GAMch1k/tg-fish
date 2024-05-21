package utils

type UserPhone struct {
	Phone string `json:"phone"`
}

type UserPhoneCode struct {
	Phone string `json:"phone"`
	Code string `json:"code"`
}

type UserPhonePassword struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type User struct {
	Phone string
	Hash string
	Code string
}