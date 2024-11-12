package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	fmt.Println(value)
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Error code: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Unknown error type")
	}
}