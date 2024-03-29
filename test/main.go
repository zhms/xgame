package main

import (
	"fmt"
	"strings"
	"xapp/xapp"
)

func main() {
	xapp.Init()
	a := "AT,BE,CY,EE,FI,FR,DE,GR,IE,IT,LV,LT,LU,MT,NL,PT,SK,SI,ES,AE"
	b := ""
	x := strings.Contains(a, b)
	fmt.Println(x)
	xapp.Run(func() {
		x := xapp.Redis().Lock("test", 0)
		x.UnLock()
	})
}
