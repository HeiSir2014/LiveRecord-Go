package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/pkg/browser"
	"net/http"
	"runtime"
)

const sListenPort = "3821";

func main() {
	fmt.Println(runtime.NumCPU(),runtime.GOOS);
	fmt.Println("Start Server Success On ")

	beego.InsertFilter("/*",beego.BeforeRouter, func(ctx *context.Context) {
		http.ServeFile(ctx.ResponseWriter,ctx.Request,"./root" + ctx.Request.URL.Path)
	})
	go RunWork();
	beego.Run(":"+sListenPort);
}

func RunWork(){
	browser.OpenURL("http://localhost:" + sListenPort + "/");
	fmt.Println("Start RunWork");
	select{}
}