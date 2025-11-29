package main

import (
	"github.com/bulka-s/wailsgen/menu"
	s "github.com/bulka-s/wailsgen/styles"
)

func main() {

	menu.Input("Project name")

	menuItems := []menu.MenuItem{
		{Text: "Web Application", Selected: false, Color: s.Blue},
		{Text: "Command Line Tool", Selected: false, Color: s.Green},
		{Text: "Library", Selected: false, Color: s.Magenta},
	}
	menu.Select("Select project type", menuItems)

}
