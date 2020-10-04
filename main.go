package main

import "tanveershafeeprottoy.com/s3-file-upload-demo/app"

func main() {
	application := &app.App{}
	application.Init()
	application.Run()
}
