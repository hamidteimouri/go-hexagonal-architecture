package main

import (
	"fmt"
	"github.com/hamidteimouri/go-hexagonal-architecture/src/app"
)

func main() {
	fmt.Println("** Balance service is running **")

	app.StartApplication()

	fmt.Println("** Balance service has been stopped **")
}
