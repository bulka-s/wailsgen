package styles

func ToBright(color string) string {
	if color == Red {
		return BrightRed
	} else if color == Green {
		return BrightGreen
	} else if color == Yellow {
		return BrightYellow
	} else if color == Blue {
		return BrightBlue
	} else if color == Magenta {
		return BrightMagenta
	} else if color == Cyan {
		return BrightCyan
	} else if color == White {
		return BrightWhite
	}
	return color
}

// ToDark makes the color dark (dim).
func ToDark(color string) string {
	return "\033[" + color + ";2m"
}
