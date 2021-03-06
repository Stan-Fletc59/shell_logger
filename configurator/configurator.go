package configurator

import (
	"fmt"
	"github.com/nvbn/shell_logger/shell"
)

func Configure(clientPath string, sh shell.Shell) {
	if sh.InWrapper() {
		fmt.Println(sh.SetupHooks(clientPath))
	} else {
		fmt.Println(sh.SetupWrapper(clientPath))
	}
}
