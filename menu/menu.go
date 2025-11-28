package menu

import (
	"fmt"

	s "github.com/bulka-s/wailsgen/styles"
)

func Input() string {
	var result string

	PSymbol(s.BorderVertical)
	Println("Project name:", s.White, s.RhombFill, s.Cyan)
	PSymbol(s.BorderVertical, s.Cyan)
	PSymbol(s.BorderBottomLeft, s.Cyan)

	MoveCursor(-2, 3)

	fmt.Scanf("%s", &result)

	MoveCursor(-2, 0)

	Println("Project name:", s.White, s.Rhomb, s.Green)
	Println(result, s.Gray, s.BorderVertical)
	PSymbol(s.BorderVertical)

	return result
}
