package menu

import (
	"fmt"

	s "github.com/bulka-s/wailsgen/styles"
)

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

func PSymbol(symbol string, smbColor ...string) {
	_smbColor := s.Gray

	if len(smbColor) > 0 {
		_smbColor = smbColor[0]
	}

	fmt.Println(_smbColor + symbol + s.Reset)
}

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
