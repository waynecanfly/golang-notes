package parse_json

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP string
}

type Serverslice struct {
	Servers []Server
}


func ParseJson()  {
	//解析到结构体
	//var s Serverslice
	//
	//str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//err := json.Unmarshal([]byte(str), &s)
	//if err != nil {
	//    return
	//}
	//fmt.Println(s)

	//解析到interface{}里面

	//b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	//
	//var f interface{}
	//err := json.Unmarshal(b, &f)
	//if err != nil {
	//    return
	//}
	//
	//m := f.(map[string]interface{})
	//
	//for k, v := range m {
	//    switch vv := v.(type) {
	//    case string:
	//        fmt.Println(k, "is string", vv)
	//    case int:
	//        fmt.Println(k, "is int", vv)
	//    case float64:
	//        fmt.Println(k, "is float64", vv)
	//    case []interface{}:
	//        fmt.Println(k, "is an array:")
	//        for i, u := range vv{
	//            fmt.Println(i, u)
	//        }
	//    default:
	//        fmt.Println(k, "is of a type I don't know how to handle")
	//    }
	//}
	//使用simplejson
	js, err := simplejson.NewJson([]byte(`{
	"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
	    }
    }`))

	if err != nil {
		return
	}
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Println("arr ->", arr)
	fmt.Println("int ->", i)
	fmt.Println("string ->", ms)
}
