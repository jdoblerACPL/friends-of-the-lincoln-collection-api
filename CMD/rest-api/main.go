package main

import router "acpl.lib.in.us/m/internal/routes"

func main() {
	r := router.InitRouter()
	r.Run()
}
