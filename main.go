package main

import (
	// local imports
	"sytron-server/routers"
)

func main() {
	r := routers.InitRouters()
	r.Run() // listen and serve on 0.0.0.0:8080
}
