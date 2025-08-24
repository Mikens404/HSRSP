package main

import "fmt"

func main() {
	fmt.Print("Hello world!")
}

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target internal/presentation --clean --package presentation api/openapi.yaml
