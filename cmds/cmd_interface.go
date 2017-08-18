package cmds

import (
	"github.com/Liam-Williams/i18n4go/common"
)

type CommandInterface interface {
	common.PrinterInterface
	Options() common.Options
	Run() error
}
