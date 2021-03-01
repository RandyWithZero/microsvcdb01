package controllers

import (
	"github.com/astaxie/beego"
	"os"
)

type MicroSVCDB01Controller struct {
	beego.Controller
}

func (c *MicroSVCDB01Controller) Get() {
	getEnv := os.Getenv("HOSTNAME")
	msg := "--->调用svcdb01" + "[" + getEnv + "]"
	c.Ctx.WriteString(msg)
}
