package main

import "practice/generate_json"

type Server struct {
    ServerName string
    ServerIP string
}

type Serverslice struct {
    Servers []Server
}

func main() {
    generate_json.GenJson()
}