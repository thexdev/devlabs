package main

import (
	"gohexarch/internal/app"
	"log"
)

func main() {
	app := app.Bootstrap()
	log.Fatal(app.Listen(":3000"))
}
