package model

type Server struct {
	ServerName string
	ServerIP string
}

type ServerSlice struct {
	Servers []Server
}