package main

import (
	"src/routing"
)


func main() {
	router := routing.RouterSetup()
	router.Run(":8080")

}
