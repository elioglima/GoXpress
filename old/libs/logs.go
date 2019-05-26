package libs

import (
	"github.com/fatih/color"
)

type Logs struct {
	*color.Color
}

// overload
func NewLogs() *Logs {
	c := &Logs{color.New(color.FgWhite)}
	return c
}
