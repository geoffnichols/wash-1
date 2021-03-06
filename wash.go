// Source for the wash executable.
//
// To extend wash, see documentation for the 'plugin' package.
package main

import (
	"os"

	"github.com/Benchkram/errz"
	"github.com/puppetlabs/wash/cmd"
	"github.com/puppetlabs/wash/config"
)

func main() {
	errz.Fatal(config.Load(), "Failed to load Wash's config")

	os.Exit(cmd.Execute())
}
