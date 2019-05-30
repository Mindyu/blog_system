package main

import "github.com/Mindyu/blog_system/app"

func main() {

	//启动http服务
	httpApp := app.NewApp()
	defer httpApp.Destory()
	_ = httpApp.Launch()

}
