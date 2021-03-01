package controllers

import (
	"github.com/astaxie/beego"
)

type MicroSVCDB01Controller struct {
	beego.Controller
}
func (c *MicroSVCDB01Controller) Get() {
	c.Ctx.WriteString("--->调用svcdb01")
}
