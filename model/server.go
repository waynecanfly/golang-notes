package model

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP string `json:"serverIP"`
}

type ServerSlice struct {
	Servers []Server `json:"servers"`
}