package main

import (
	"github.com/astaxie/beego"
	_ "microsvcdb01/conf"
	_ "microsvcdb01/filter"
	"microsvcdb01/routers"
)

func main() {
	routers.Router()
	beego.Run()

}
