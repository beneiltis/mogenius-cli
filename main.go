/*
Copyright © 2022 mogenius GmbH https://mogenius.com

*/
package main

import (
	_ "embed"
	"fmt"

	"github.com/beneiltis/mogenius-cli/cmd"
)

//go:generate bash get_version.sh
//go:embed version.txt
var version string

func main() {
	fmt.Printf("Version: %s\n\n", version)
	cmd.Execute()
}
