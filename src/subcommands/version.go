package subcommands

import (
	"embed"
	"fmt"
	"os"

	"github.com/agu3rra/odincli/src/common"
)

//go:embed version
var versionFile embed.FS

func Version() {

	version, err := versionFile.ReadFile("version")
	if err != nil {
		common.LogErr.Printf("Error reading version information: %v\n", err)
		os.Exit(common.ExitRuntimeError)
	}
	fmt.Printf("odin version: %s\n", string(version))
	os.Exit(common.ExitOK)
}
