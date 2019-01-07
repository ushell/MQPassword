package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/ushell/MQPassword"
)

var mainwin *ui.Window

func main() {
	ui.Main(MQPassword.InitUI)
}
