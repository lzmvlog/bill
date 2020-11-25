package main

import "bill/routers"

func main()  {
	// 注册路由
	r := routers.Router()
	r.Run()
}
