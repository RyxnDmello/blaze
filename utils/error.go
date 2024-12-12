package utils

import (
	"fmt"
	"os"
)

func Error(err interface{}) {
	fmt.Printf("\n\n󰈸 Blaze CLI ::: %s%s%s\n\n", "\033[31m", err.(string), "\033[0m")
	os.Exit(0)
}
