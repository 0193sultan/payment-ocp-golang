package main

import (
	"fmt"

	"github.com/0193sultan/payment-ocp-golang/internal/handler/http"
)

func main() {
	fmt.Println("OCP Payment Service Running...")

	r := http.SetupRouter()
	r.Run(":8080")
}
