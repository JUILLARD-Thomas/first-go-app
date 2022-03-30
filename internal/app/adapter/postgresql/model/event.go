package model

type Event struct {
	Id         int    `json:"id"`
	Type_enum  string `json:"type"`
	User_agent string `json:"user_agent"`
	Ip         string `json:"ip"`
	Ts         string `json:"timestamp"`
}
