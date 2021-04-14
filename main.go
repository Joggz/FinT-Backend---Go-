package main

import (
	"fmt"
	"github.com/Joggz/FintT-Backend---Go-.git/Migrations"
)

func main(){
	fmt.Println("running")
	Migrations.Migrate()
}
