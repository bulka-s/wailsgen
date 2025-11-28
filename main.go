package main

import (
	"fmt"

	"github.com/bulka-s/wails-ts/styles"
)

func main() {
	fmt.Println(styles.Bold + styles.Green + "Hello, Wails!" + styles.Reset)
}
