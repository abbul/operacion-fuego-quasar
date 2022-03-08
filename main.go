package main

import "github.com/abbul/operacion-fuego-quasar/app"

func main() {
	var engine = app.StartApp()
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
