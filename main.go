package main

import (
	"github.com/justyntemme/rockemsockemrobots/gameboard"
	"github.com/justyntemme/rockemsockemrobots/web"
)

func main() {
	web.StartServer()
	gameboard.Init()
}
