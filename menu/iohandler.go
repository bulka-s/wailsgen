package menu

import (
	"fmt"

	s "github.com/bulka-s/wailsgen/styles"
)

type MenuItem struct {
	Text     string
	Selected bool
	Color    string
}

// Print writes text without a newline.
// textColor is an ANSI color code from the styles package.
// symbol is an optional leading symbol printed before the text.
// smbColor (optional) sets the color for the symbol.
func Print(text string, textColor string, symbol string, smbColor ...string) {
	_smbColor := s.Gray

	if len(smbColor) > 0 {
		_smbColor = smbColor[0]
	}

	if symbol != "" {
		fmt.Printf("%s%s%s  %s%s%s", _smbColor, symbol, s.Reset, textColor, text, s.Reset)
	} else {
		fmt.Printf("%s%s%s", textColor, text, s.Reset)
	}
}

// Println writes text followed by a newline.
// Parameters are the same as Print.
func Println(text string, textColor string, symbol string, smbColor ...string) {
	_smbColor := s.Gray

	if len(smbColor) > 0 {
		_smbColor = smbColor[0]
	}

	if symbol != "" {
		fmt.Printf("%s%s%s  %s%s%s\n", _smbColor, symbol, s.Reset, textColor, text, s.Reset)
	} else {
		fmt.Printf("%s%s%s\n", textColor, text, s.Reset)
	}
}

// PSymbol prints a single symbol on its own line.
// smbColor (optional) sets the color for the symbol.
func PSymbol(symbol string, smbColor ...string) {
	_smbColor := s.Gray

	if len(smbColor) > 0 {
		_smbColor = smbColor[0]
	}

	fmt.Println(_smbColor + symbol + s.Reset)
}

func PrintMenuItem(text string, isSelected bool, textColor string) {
	bullet := s.BulletUnselected
	if isSelected {
		bullet = s.BulletSelected
	}

	color := textColor
	if !isSelected {
		color = s.ToBright(textColor)
	}

	fmt.Printf("%s%s%s  %s%s %s%s\n",
		s.Gray, s.BorderVertical, s.Reset,
		color, bullet, text, s.Reset)
}

// MoveCursor moves the terminal cursor by rows and cols.
// Negative rows move the cursor up, positive rows move it down.
// Negative cols move the cursor left, positive cols move it right.
func MoveCursor(rows, cols int) {
	if rows < 0 {
		fmt.Printf("\033[%dA", -rows)
	} else if rows > 0 {
		fmt.Printf("\033[%dB", rows)
	}
	if cols < 0 {
		fmt.Printf("\033[%dD", -cols)
	} else if cols > 0 {
		fmt.Printf("\033[%dC", cols)
	}
}
