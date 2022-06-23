package main

import (
	"cloud_homestay/routers"
)

func main() {
	r := routers.Router()
	r.Run()
}
