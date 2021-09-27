package generate_json

import (
	"encoding/json"
	"fmt"
	"practice/model"
)

//生成json
//假设我们需要生成服务器列表信息，


func GenJson()  {
	var s model.ServerSlice
	s.Servers = append(s.Servers, model.Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, model.Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})

	b, err := json.Marshal(s)

	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(string(b))
}