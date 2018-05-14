package main

import (
	"dingding/services/gateway/api"
	"dingding/services"
)

func main() {
	services.RunService(api.NewService())
}
