package main

import (
	"github.com/Aj002Th/LittlePrince/router"
	"github.com/Aj002Th/LittlePrince/service"
)

func main() {
	service.Setup()
	router.Setup()
}
