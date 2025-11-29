package menu

import (
	"bufio"
	"os"
	"strings"

	s "github.com/bulka-s/wailsgen/styles"
	"github.com/eiannone/keyboard"
)

func Input(title string) string {
	var result string

	PSymbol(s.BorderVertical)
	Println(title, s.White, s.RhombFill, s.Cyan)
	PSymbol(s.BorderVertical, s.Cyan)
	PSymbol(s.BorderBottomLeft, s.Cyan)

	MoveCursor(-2, 3)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	result = strings.TrimSpace(line)

	MoveCursor(-2, 0)

	Println(title, s.White, s.Rhomb, s.Green)
	Println(result, s.Gray, s.BorderVertical)

	return result
}

func Select(title string, options []MenuItem) string {
	var result string
	var selectedIndex int = 0

	keyboard.Open()
	defer keyboard.Close()

	PSymbol(s.BorderVertical)
	Println(title, s.White, s.RhombFill, s.Cyan)

	for {
		for _, option := range options {
			PrintMenuItem(option.Text, option.Selected, option.Color)
		}

		_, key, _ := keyboard.GetKey()

		if key == keyboard.KeyArrowUp {
			if selectedIndex > 0 {
				options[selectedIndex].Selected = false
				selectedIndex--
				options[selectedIndex].Selected = true
			} else {
				options[selectedIndex].Selected = false
				selectedIndex = len(options) - 1
				options[selectedIndex].Selected = true
			}
		} else if key == keyboard.KeyArrowDown {
			if selectedIndex < len(options)-1 {
				options[selectedIndex].Selected = false
				selectedIndex++
				options[selectedIndex].Selected = true
			} else {
				options[selectedIndex].Selected = false
				selectedIndex = 0
				options[selectedIndex].Selected = true
			}
		} else if key == keyboard.KeyEnter {
			result = options[selectedIndex].Text
			break
		}
		MoveCursor(-len(options), 0)
	}

	return result
}
