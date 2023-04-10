package main

import "challenge-chapter-2-sesi-2/routers"

func main() {
	var PORT = ":80"

	routers.Route().Run(PORT)
}
