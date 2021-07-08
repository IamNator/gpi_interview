package main

import (
	"fmt"

	"github.com/IamNator/gpi_interview/buyandsell"
)

func main() {
	var cement buyandsell.Cement
	cement.BuyCement(15)
	cement.SellCement(9)
	fmt.Println(cement)
}
