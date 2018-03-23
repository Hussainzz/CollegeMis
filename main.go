package main

import (
	_ "CollegeMis/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
)
func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	  }))	
	
	  orm.RegisterDriver("mysql", orm.DRMySQL)
	  orm.RegisterDataBase("default",  "mysql",  "root:admin@/misapi?charset=utf8")
	  name := "default"
	  force := false
	  verbose := false
	  err := orm.RunSyncdb(name, force, verbose)
	  if err != nil{
		  fmt.Println(err)
	  }
	}


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
