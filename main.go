package main

import (
	"Blog/server"
)

func main()  {
	App := server.New()
	App.Run(":8081")
}
